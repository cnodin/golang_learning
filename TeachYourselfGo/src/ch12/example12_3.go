package main

import (
	"fmt"
	"time"
)

func receiver(c chan string) {
	for s := range c {
		fmt.Println(s)
	}
}

func main() {
	messages := make(chan string, 2)
	messages <- "hello"
	messages <- "world"
	close(messages)

	fmt.Println("Pushed two messages onto Channel with no receivers")
	time.Sleep(time.Second * 1)
	receiver(messages)
}
