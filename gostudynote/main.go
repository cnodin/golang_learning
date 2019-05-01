package main

import (
	"fmt"
)

func main() {
	/*println("hello, go")
	println("hello, go")
	fmt.Println("hello, world")
	fmt.Println("hello, world")

	var x int32
	var s = "hello, world"

	println(x, s)

	y := 100
	println(x, s, y)*/

	/*x := 100

	if (x > 0) {
		println("x")
	}else if (x < 0) {
		println("-x")
	}else {
		println("0")
	}

	switch {
	case x > 0 :
		println("x")
	case x < 0 :
		println("-x")
	default :
		println("0")
	}

	for i := 0; i < 5; i++ {
		println(i)
	}

	for i := 4; i >= 0; i-- {
		println(i)
	}

	for x < 5 {
		println(x)
		x++
	}

	for {
		println(x)
		x--

		if x < 0 {
			break;
		}
	}

	x1 := []int{100, 101, 102}

	for i, n := range x1 {
		println(i, ":", n)
	}*/

	/*a, b := 10, 0
	c, err := div(a, b)

	fmt.Println(c, err)

	f := test(a);
	f()*/

	//定义切片，int类型，初始容量5个
	/*
	x := make([]int, 0, 5)

	for i := 0; i < 8; i++ {
		x = append(x, i)
	}
	fmt.Println(x)
	*/

	//定义map类型，key是string，value是int
/*	m := make(map[string]int)

	m["a"] = 1

	//使用ok-idom模式。如果m["a"]存在，则ok返回true；否则返回false
	x, ok := m["a"]
	fmt.Println(x, ok)

	delete(m, "a")*/

	var m manager

	m.name = "Tom"
	m.age = 29
	m.title = "CTO"

	fmt.Println(m)
}

type user struct {
	name string
	age byte
}

type manager struct {
	user
	title string
}

func div(a, b int) (int, error) {
	//if b == 0 {
	//	return 0, errors.New("division by zero")
	//}
	defer println("dispose...")

	return a / b, nil
}

func test(x int) func() {
	return func(){
		println(x)
	}
}
