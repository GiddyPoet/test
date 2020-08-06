package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func handleConnect(location, addr string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return
	}
	buf := make([]byte, 256)

	for {
		conn.Read(buf)
		fmt.Print(location + ": " + string(buf))
	}
}

func main() {
	fmt.Println("args nums", len(os.Args))
	if len(os.Args) < 1 {
		fmt.Println("usage: location=ip:port")
	}
	for _, value := range os.Args[1:] {
		str := strings.Split(value, "=")
		go handleConnect(str[0], str[1])
	}
	for {

	}
}
