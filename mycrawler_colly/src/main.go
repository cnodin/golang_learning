package src

import (
	"./model"
	"encoding/json"
	"github.com/gocolly/colly"
	"log"
)

func main() {
	c := colly.NewCollector()

	loginRequest := &model.LoginRequest{
		AppID: 1,
		Captcha: "",
		Cellphone: "13599521190",
		Country: 86,
		Password: "",
		Platform: 3,
		Remember: 1,
	}

	bData, _:= json.MarshalIndent(loginRequest, "", "\t")

	//model.LoginRequest{}
	err := c.PostRaw("https://account.geekbang.org/account/ticket/login", bData);
	if err != nil {
		log.Fatalln(err)
	}
}
