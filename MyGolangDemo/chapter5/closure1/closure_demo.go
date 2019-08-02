package main

import "fmt"

func main() {
	nextNumber := getSequence()

	fmt.Printf("%d ", nextNumber())
	fmt.Printf("%d ", nextNumber())
	fmt.Printf("%d ", nextNumber())

	//创建新的函数nextNumber1，并查看结果
	nextNumber1 := getSequence()
	fmt.Printf("%d ", nextNumber1())
	fmt.Printf("%d ", nextNumber1())
}

func getSequence() func() int {
	i := 0
	return func() int {
		i++

		return i
	}
}
