package main

import (
	_ "log"
	"runtime/debug"
)

func test() {
	println("i am dead")
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
		}
		debug.PrintStack()
	}()

	test()
}
