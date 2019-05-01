package main

import "fmt"

func test1(s string, a ...int) {
	fmt.Printf("%T, %v\n", a, a) //显示类型和值
}

func test2(a ...int) {
	for i := range a {
		fmt.Printf("%T, %v\n", i, i)
	}
}

func test3(a ...int) {
	for i := range a {
		a[i] += 100	//会修改底层数据的值
	}
}

func main() {
	test1("abc", 1, 2, 3, 4)
	a := [3]int{1,2,3}
	test1("efg", a[:]...) //转换为slice后展开
	test2(a[:]...)
	test3(a[:]...)
	fmt.Println(a)
}
