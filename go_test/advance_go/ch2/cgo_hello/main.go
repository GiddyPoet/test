package main

import "C" // cgo

//#include <stdio.h>
func main() {
	println("hello cgo")
	C.puts(C.CString("hello, world\n"))
}
