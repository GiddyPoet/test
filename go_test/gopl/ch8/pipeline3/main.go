package main

import "fmt"

func counter(out chan<- int) {
	for i := 0; i < 100; i++ {
		out <- i
	}
	close(out)
}

func double(in <-chan int, out chan<- int) {
	for val := range in {
		out <- val * val
	}
	close(out)
}

func show(in <-chan int) {
	for val := range in {
		fmt.Println(val)
	}
}

func main() {
	line1 := make(chan int)
	line2 := make(chan int)

	go counter(line1)
	go double(line1, line2)
	show(line2)
}
