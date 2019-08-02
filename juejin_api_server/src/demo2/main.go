package main

import (
	"demo2/config"
	"demo2/router"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"time"
)

var (
	cfg = pflag.StringP("config", "c", "", "apiserver config file path.")
)

const (
	RUNMODE        = "runmode"
	MAX_PING_COUNT = "max_ping_count"
	URL            = "url"
	ADDR           = "addr"
)

func main() {
	pflag.Parse()

	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	//set gin mode
	gin.SetMode(viper.GetString(RUNMODE))

	g := gin.New()
	moddleware := []gin.HandlerFunc{}
	router.Load(g, moddleware...)

	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up", err)
		}
		log.Print("The router has been deploy successfully.")
	}()

	log.Printf("Start to listening the incoming requests on http address: %s", viper.GetString(ADDR))
	log.Printf(http.ListenAndServe(viper.GetString(ADDR), g).Error())
}

func pingServer() error {
	for i := 0; i < viper.GetInt(MAX_PING_COUNT); i++ {
		resp, err := http.Get(viper.GetString(URL) + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping.
		log.Print("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}

	return errors.New("Cannot connect to the router.")
}
