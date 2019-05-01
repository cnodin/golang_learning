package main

import (
	"os"
	"fmt"
	"net/http"
	"io"
)

func init() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./list35 <url>")
	}
}

func main() {
	r, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	io.Copy(os.Stdout, r.Body)
	if err := r.Body.Close(); err != nil {
		fmt.Println(err)
	}


}
