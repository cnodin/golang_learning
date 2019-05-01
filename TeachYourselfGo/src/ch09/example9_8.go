package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	var buffer bytes.Buffer

	for i := 0; i < 500; i++ {
		buffer.WriteString("z")
	}

	fmt.Println(buffer.String())
	fmt.Println(len(buffer.String()))
	fmt.Println(strings.ToUpper(buffer.String()))
}
