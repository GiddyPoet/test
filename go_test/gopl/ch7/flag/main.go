package main

import (
	"flag"
	"fmt"
	"time"
)

// flag.Duration is time param name default param help message
var period = flag.Duration("period", 1*time.Second, "sleep period")

func main() {
	flag.Parse()
	fmt.Printf("Sleeping for %v", *period)
	time.Sleep(*period)
	fmt.Println()
}
