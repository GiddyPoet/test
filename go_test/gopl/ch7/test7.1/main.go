package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type byteCounter int
type wordCounter int
type lineCounter int

// SplitFunc is fun type
// type SplitFunc func(data []byte, atEOF bool) (advance int, token []byte, err error)
func retCount(p []byte, fn bufio.SplitFunc) int {
	s := string(p)
	// func NewScanner(r io.Reader) *Scanner
	// func NewReader(s string) *Reader return the io.Reader
	scanner := bufio.NewScanner(strings.NewReader(s))
	// func (s *Scanner) Split(split SplitFunc)
	// set scan split function split the string
	scanner.Split(fn)
	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "reading input:%v", err)
	}
	return count
}

func (c *byteCounter) Write(p []byte) (int, error) {
	*c += byteCounter(len(p))
	return len(p), nil
}

func (c *wordCounter) Write(p []byte) (int, error) {
	count := retCount(p, bufio.ScanWords)
	*c += wordCounter(count)
	return count, nil
}

func (c *lineCounter) Write(p []byte) (int, error) {
	count := retCount(p, bufio.ScanLines)
	*c += lineCounter(count)
	return count, nil
}

//type Writer interface{
// Write(p []byte) (n int, err error)
// }
// 只要实例化了上述接口，任何约定为io.Writer的入参都可以传入，因为其满足了接口的约定
func main() {
	var c byteCounter
	c.Write([]byte("Hello this is a line"))
	fmt.Println("Byte Counter ", c)

	var w wordCounter
	w.Write([]byte("Hello this is a line"))
	fmt.Println("word Counter", w)

	var l lineCounter
	l.Write([]byte("Hello this is a line"))
	fmt.Println("line Counter", l)

	c, w, l = 0, 0, 0
	test := "GiddyPoet"
	fmt.Fprintf(&c, "Hello this is a line %s", test)
	fmt.Println(c)

	fmt.Fprintf(&w, "Hello this is a line %s", test)
	fmt.Println(w)

	fmt.Fprintf(&l, "Hello this is a line %s", test)
	fmt.Println(l)
}
