package main

import "fmt"

func main() {
	line1 := make(chan int)
	line2 := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			line1 <- i
		}
		close(line1)
	}()

	go func() {
		for val := range line1 {
			line2 <- val * val
		}
		close(line2)
	}()

	for x := range line2 {
		fmt.Println(x)
	}
}
