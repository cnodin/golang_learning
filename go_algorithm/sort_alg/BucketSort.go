package main

import "fmt"

func bucket_sort(length int, array []int) {
	bucket := make([]int, length)
	for i := 0; i < len(array); i++ {
		bucket[array[i]]++
	}

	//for i := len(bucket) - 1; i >= 0; i-- {
	//	for j := 0; j < bucket[i]; j++ {
	//		fmt.Println(i)
	//	}
	//}

	for i := 0; i < len(bucket); i++ {
		for j := 0; j < bucket[i]; j++ {
			fmt.Println(i)
		}
	}
}

func main() {
	array := []int{5, 10000000000, 1, 9, 5, 3, 7, 6, 1}
	bucket_sort(10000000001, array)
}
