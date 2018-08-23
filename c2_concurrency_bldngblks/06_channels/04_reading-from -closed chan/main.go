package main

import (
	"fmt"
)

func main() {
	intStream := make(chan int)
	go func() {
		intStream <- 5
	}()
	integer, ok := <-intStream
	fmt.Printf("(%v): %v", ok, integer)
}
