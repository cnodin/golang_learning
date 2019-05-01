package main

import (
	"config"
	"fmt"
	"github.com/siddontang/go-mysql/client"
)

func main() {
	conn, err := client.Connect("127.0.0.1:3306", "root", "rootroot", "test")
	if err != nil {
		fmt.Println("has error: " + err.Error())
		return
	}
	conn.Ping()
	config.LoadConfig()
	fmt.Println("Hello go_mysql_client!")
}
