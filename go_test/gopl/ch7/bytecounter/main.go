package main

import "fmt"

type ByteCounter int

// the Write interface implement for go test
func (b *ByteCounter) Write(p []byte) (int, error) {
	*b += ByteCounter(len(p)) // b is pointer so *b is b value
	return len(p), nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))

	fmt.Println(c)

	c = 0
	var name = "GiddyPoet"
	// Fprintf in put io.Write is a interface , if you inmplemet this you can use use this struct for function
	fmt.Fprintf(&c, "hello , %s", name)
	fmt.Println(c)
}
