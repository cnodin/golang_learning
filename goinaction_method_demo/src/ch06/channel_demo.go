package main

import "fmt"

func main() {
	buffered := make(chan string, 10)

	buffered <- "Gopher"

	value := <-buffered

	fmt.Println(value)
}
