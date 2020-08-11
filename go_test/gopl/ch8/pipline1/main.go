package main

import "fmt"

func main() {
	line1 := make(chan int)
	line2 := make(chan int)

	go func() {
		for x := 0; ; x++ {
			line1 <- x
		}
	}()

	go func() {
		for {
			x := <-line1
			line2 <- x * x
		}
	}()

	for {
		fmt.Println(<-line2)
	}
}
