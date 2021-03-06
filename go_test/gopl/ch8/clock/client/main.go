package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	// the conn implement read write close and other interface so it can be io.Writer io.Reader
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	// copy the conn stream to stdout stream
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
