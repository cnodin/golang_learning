package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	var b = true
	fmt.Println(reflect.TypeOf(b))
	var s = strconv.FormatBool(true)
	var s1 = strconv.QuoteToGraphic(s)
	fmt.Println(reflect.TypeOf(s))
	fmt.Println(s1)
	fmt.Println(s)
}
