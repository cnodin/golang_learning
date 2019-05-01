package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello")

	//array := [...]float64{7.0, 8.5, 9.1}
	//x := Sum(&array)
	//
	//fmt.Printf("x: ", x)
	pid := os.Getpid();
	fmt.Println("pid: ", pid)

	ppid := os.Getppid();
	fmt.Println("ppid: ", ppid)
}

func Sum(a *[3]float64) (sum float64) {
	for _, v := range *a {
		sum += v
	}

	return sum
}
