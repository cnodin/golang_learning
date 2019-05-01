package main

import "fmt"

func sort(array []int) {
	length := len(array)
	if (length > 0) {
		for i := 1; i < length; i++ {
			for j := 0; j < length - i; j++ {
				if (array[j] > array[j + 1]) {
					temp := array[j]
					array[j] = array[j+1]
					array[j+1] = temp
				}
			}
		}
	}
}

func sort1(array []int) {
	length := len(array)
	if (length > 0) {
		for i := 0; i < length - 1; i++ {
			for j := i+1; j < length; j++ {
				if (array[i] > array[j]) {
					array[i], array[j] = array[j], array[i]
				}
			}
		}
	}
}

func main() {
	array := []int{2,4,5,1,6}
	sort1(array)
	for i, _ := range array {
		fmt.Println(array[i])
	}
}
