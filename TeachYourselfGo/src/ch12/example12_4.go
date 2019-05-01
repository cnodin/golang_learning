package main

import (
	"time"
	"fmt"
)

func pinger(c chan string) {
	t := time.NewTicker(1 * time.Second)
	for {
		c <- "ping"
		<- t.C
	}
}

func channelReader(messages <-chan string) {
	msg := <-messages
	fmt.Println("msg in channelReader:" + msg)
}

func channelWriter(messages chan<- string) {
	messages <- "Hello world"
}

func channelReaderAndWriter(messages chan string) {
	msg := <-messages
	fmt.Println(msg)
	messages <- "hello world"
}

func main() {

	messages := make(chan string)
	go pinger(messages)
	for {
		msg := <-messages
		fmt.Println(msg)
		channelReader(messages)
	}
}
