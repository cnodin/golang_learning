package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	message := "Away from keyboard. https://golang.com"

	encodingMessage := base64.StdEncoding.EncodeToString([]byte(message))

	fmt.Println(encodingMessage)

	data, err := base64.StdEncoding.DecodeString(encodingMessage)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(data))
	}

}
