package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("my server starting...")

	listener, err := net.Listen("tcp", "127.0.0.1:8085")
	conn, err := listener.Accept()


}
