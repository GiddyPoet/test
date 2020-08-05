package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func test() {
	test := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		test[input.Text()]++
	}

	for line, n := range test {
		fmt.Printf("%s\t%d", line, n)
	}
}

func testLog() {
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		log.Println(input.Text())
	}
}

func testString() {
	// create a Reader implement
	r := strings.NewReader("ABC\nDEF\r\nGHI\nJKL")
	input := bufio.NewScanner(r)
	for input.Scan() {
		fmt.Printf("%d %v\n", input.Bytes(), input.Text())
	}
}

func testWrite() {
	w := bufio.NewWriter(os.Stdout)
	fmt.Fprint(w, "Hello, ")
	fmt.Fprint(w, "world!\n")
	// if don't flush the data buffer will not send to io.Write , so we can't get the message form stdin
	w.Flush()
}
func testWriteString() {
	input := "1234 5678 12345678901234567890"
	// strings.NewReader return a io.Reader , NewScanner just need a io.Reader
	scanner := bufio.NewScanner(strings.NewReader(input))
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		advance, token, err = bufio.ScanWords(data, atEOF)
		if err == nil && token != nil {
			_, err = strconv.ParseInt(string(token), 10, 32)
		}
		return
	}
	scanner.Split(split)

	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid input: %s", err)
	}
}

func main() {
	testWrite()
}
