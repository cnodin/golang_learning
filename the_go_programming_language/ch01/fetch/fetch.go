package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		fetchUrl(url)
	}
}

func fetchUrl(url string) {
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}
	resp, err := http.Get(url)
	//defer resp.Body.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		return
	}

	//b, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode == 200 {
		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		resp.Body.Close()
	} else {
		fmt.Printf("fetch error, http status code: %d", resp.StatusCode)
	}
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
	//	return
	//}
	//fmt.Printf("%s", b)
}
