package main

import (
	"log"
	"os"

	_ "./matchers"
	_ "./search"
	"fmt"
)

func main() {
	//search.Run("president")
	//slice := []int{10, 20, 30, 40, 50}
	//newSlice := slice[1:3]
	//newSlice[1] = 35
	//newSlice1 := append(slice, 60)

	//source := [4]int{10, 20, 30, 40}
	//slice := source[2:3]
	//slice[0] = 35
	//slice = append(slice, 50, 60, 70, 80)

	//for index, value := range slice {
	//	fmt.Printf("Index: %d Value: %d\n", index, value)
	//}
	//fmt.Printf("%v\n", slice)

	colors := map[string]string{
		"AliceBlue": "#f0f8ff",
		"Coral":     "#ff7F50",
		"DarkGray":  "#a9a9a9",
	}

	displayColors(colors)
	removeColor(colors, "Coral")
	displayColors(colors)
}

/**
显示颜色
*/
func displayColors(colors map[string]string) {
	for key, value := range colors {
		fmt.Printf("Key: %s Value: %s\n", key, value)
	}
}

/**
移除颜色
*/
func removeColor(colors map[string]string, key string) {
	delete(colors, key)
}

func init() {
	log.SetOutput(os.Stdout)
}
