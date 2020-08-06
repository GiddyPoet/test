package main

import (
	"io"
	"log"
	"net"
	"os"
)

func mustCopy(dst io.Writer, src io.Reader) {
	io.Copy(dst, src)
}

func main() {
	conn, err := net.Dial("tcp", "0.0.0.0:7353")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()
	// if the conn have msg output to os.stdout
	go mustCopy(os.Stdout, conn)
	// send input to server
	mustCopy(conn, os.Stdin)
}
