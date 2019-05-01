package main

import (
	"fmt"
	"sync"
)

func main() {
	m := map[string]string {
		"W": "forward",
		"A": "left",
		"D": "right",
		"S": "backward",
	}

	for k,v := range m {
		fmt.Println(k, v)
	}
}
