package main

import "fmt"

func echo(s string) {
	fmt.Println(s)
}

func main() {
	s := "hello world"
	t := "Goodbye Cruel World"

	echo(s)
	echo(t)
}
