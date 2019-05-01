package main

import (
	"fmt"
)

type NewInt int

type IntAlias = int

//type MyDuration = time.Duration

//func (m MyDuration) EasySet(a string) {
//	fmt.Println(a)
//}

func (n IntAlias) EasySet(a string) {
	fmt.Println(a)
}

func main() {

	var a NewInt
	fmt.Printf("a type: %T\n", a)

	var a2 IntAlias
	fmt.Printf("a2 type: %T\n", a2)

	//var d MyDuration
	//d.EasySet("aaa")
	a.EasySet("bb")
}
