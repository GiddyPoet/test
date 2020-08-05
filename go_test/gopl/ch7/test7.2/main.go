// this maybe not right

package main

import (
	"fmt"
	"io"
	"os"
)

type CounteWriter struct {
	cw io.Writer
	C  int64
}

// interface implement
func (c *CounteWriter) Write(p []byte) (int, error) {
	c.C = int64(len(p))
	n, err := c.cw.Write(p)
	return n, err
}

// CountingWriter *int64 return the value with the msg change
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	c := &CounteWriter{cw: w}
	return c, &c.C
}

func main() {
	rw, line := CountingWriter(os.Stdout)
	rw.Write([]byte("Hello, world\n"))
	fmt.Println(*line)
}
