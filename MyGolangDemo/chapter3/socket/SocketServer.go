package main

import (
	"net"
	"strings"
	"fmt"
	"time"
	"io"
	"bytes"
)

const (
	SERVER_NETWORK = "tcp"
	SERVER_ADDRESS = "127.0.0.1:8085"
	DELIMITER = '\t'
)
var listener net.Listener

func main() {
}

func init()  {
	listener, err := net.Listen(SERVER_NETWORK, SERVER_ADDRESS)
	if err != nil {
		printServerLog("Listen Error: %s", err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			printServerLog("Accept Error: %s", err)
			return
		}
		printServerLog("Established a connection with a client application. " +
			"(remote address: %s", conn.RemoteAddr())

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn)  {
	for {
		conn.SetReadDeadline(time.Now().Add(10 * time.Second))
		strReq, err := read(conn)
		if err != nil {
			if err == io.EOF {
				printServerLog("The connection is closed by another side.")
			} else {
				printServerLog("Read Error: %s", err)
			}
			break
		}
		printServerLog("Received request: %s", strReq)
	}
}
func read(conn net.Conn) (string, error) {
	readBytes := make([]byte, 1)
	var buffer bytes.Buffer

	for {
		_, err := conn.Read(readBytes)
		if err != nil {
			return "", err
		}
		readByte := readBytes[0]
		if readByte == DELIMITER {
			break
		}
		buffer.WriteByte(readByte)
	}

	return buffer.String(), nil
}

func printLog(role string, sn int, format string, args ...interface{})  {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	fmt.Printf("%s[%d]: %s", role, sn, fmt.Sprintf(format, args...))
}

func printServerLog(format string, args ...interface{})  {
	printLog("Server", 0, format, args...)
}