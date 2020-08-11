package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func test() {
	done := make(chan int)
	go func() {
		fmt.Println("hellow world")
		<-done
	}()
	// time.Sleep(1 * time.Second)
	done <- 1
}

func test1() {
	var mu sync.Mutex

	mu.Lock()
	go func() {
		fmt.Println("hello world")
		mu.Unlock()
	}()
	mu.Lock()
}

func test2() {
	tick := time.Tick(10 * time.Microsecond)
	go func() {
		fmt.Println("hello world")
	}()
	<-tick
}

func test3() {
	done := make(chan int, 1)
	// buf channel
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("hello world")
			done <- 1
		}()
	}

	for i := 0; i < 10; i++ {
		<-done
	}
}

func test4() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		// add wait event num
		wg.Add(1)
		go func() {
			fmt.Println("hello world")
			// done a wait event num --
			wg.Done()
		}()
	}
	// wait all event done
	wg.Wait()
}

// Producer product
func Producer(factor int, out chan<- int) {
	for i := 0; ; i++ {
		out <- i * factor
	}
}

// Consumer consum
func Consumer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func test5() {
	var ch = make(chan int, 64)
	go Producer(1, ch)
	go Producer(4, ch)

	go Consumer(ch)

	time.Sleep(5 * time.Second)
}

func test6() {
	ch := make(chan int, 64)

	go Producer(1, ch)
	go Producer(4, ch)
	go Consumer(ch)
	sig := make(chan os.Signal)

	// signal 
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("quit (%v)\n", <-sig)
}
func main() {
	test6()
}
