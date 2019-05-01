package main

import (
	"time"
	"fmt"
)

func slowFunc() {
	time.Sleep(time.Second * 5)
	fmt.Println("sleeper() finished")
}

func main() {
	go slowFunc()
	fmt.Println("i am now shown straightaway")
	time.Sleep(time.Second * 6)
}
