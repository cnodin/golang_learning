package main

import "fmt"

func main() {
	fmt.Printf("%d\n", Add()(3))
	fmt.Printf("%d\n", Add2(6)(3))
}

func Add() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func Add2(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}
