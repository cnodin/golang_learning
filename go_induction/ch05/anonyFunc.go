package main

import "fmt"

func visit(list []int, f func(int))  {
	for _, v := range list {
		f(v)
	}
}

func main() {

	visit([]int{1, 2, 3}, func(i int) {
		fmt.Println(i)
	})

}
