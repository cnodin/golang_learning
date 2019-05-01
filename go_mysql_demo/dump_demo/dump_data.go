package main

import (
	"github.com/siddontang/go-mysql/client"
	"fmt"
)

func main() {
	//conn, _ := client
	conn, _ := client.Connect("127.0.0.1:3306", "root", "rootroot", "test")
	conn.Ping()

	r, _ := conn.Execute("SELECT * FROM PERSON_INFO")
	fmt.Println(r)
	fmt.Println("finished")
}
