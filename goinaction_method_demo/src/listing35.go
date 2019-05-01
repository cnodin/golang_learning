package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var b bytes.Buffer

	//write string to byte
	b.Write([]byte("hello"))

	fmt.Fprintf(&b, " World")

	io.Copy(os.Stdout, &b)
}
