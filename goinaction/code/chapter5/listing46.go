package main

import "fmt"

type duration int

func (d duration) pretry() {
	fmt.Printf("Duration: %d\n", &d)
}

func main() {
	duration(42).pretry()
}
