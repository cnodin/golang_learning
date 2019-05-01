package main

import (
	"net/http"
	"fmt"
	_"strings"
	"log"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	/*r.ParseForm() //解析参数
	fmt.Println(r.Form) //在服务端打印参数
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v," "))
	}*/

	fmt.Fprintf(w, "hello, goweb")
}

func main() {

	http.HandleFunc("/", sayhelloName) //设置访问的路由
	http.Serve()
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}


}
