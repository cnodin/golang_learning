package main

import (
	"github.com/hashicorp/consul/api"
	"fmt"
)

func main() {
	client, err := api.NewClient(&api.Config{Address: "127.0.0.1:8500"})
	if err != nil {
		fmt.Println("new client has error: ", err)
	}

	client.LockKey("webhook_receiver/1")

}
