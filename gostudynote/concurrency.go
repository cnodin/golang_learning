package main

import (
	"fmt"
	"time"
)

func main() {
/*	go task(1)
	go task(2)

	time.Sleep(time.Second * 6)*/

	var a = 10
	b := 20

	println(a)
	println(b)

	done := make(chan bool)
	data := make(chan int)

	go consumer(data, done)
	go producer(data)

	<-done
}


func task(id int)  {
	for i := 0; i < 5; i++  {
		fmt.Printf("%d: %d\n", id, i)
		time.Sleep(time.Second)
	}
}

func consumer(data chan int, done chan bool) {
	for x := range data {
		println("recv:", x)
	}

	done <- true
}

func producer(data chan int) {
	for i := 0; i < 4; i++ {
		data <- i
	}

	close(data)
}
