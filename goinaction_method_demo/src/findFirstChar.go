package main

import "fmt"

func main() {
	str := "aabbcddeffg"
	findFirstChar(str)
}

func findFirstChar(str string)  {
	theMap := make(map[int32]int)
	for _, v := range str {
		fmt.Printf("%c\t", v)
		theMap[v]++
	}

	for key, value := range theMap {
		fmt.Printf("\nKey: %d	Value: %d", key, value)
	}
}
