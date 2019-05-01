package main

import (
	"log"
	"os"
	_ "./matchers"
	_ "./search"
	"golang.org/x/text/search"
)

//init函数在main之前调用
func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
	//fmt.Println("hello")
	println(Matcher)
}
