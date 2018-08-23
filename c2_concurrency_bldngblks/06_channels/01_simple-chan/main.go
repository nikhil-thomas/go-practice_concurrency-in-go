package main

import (
	"fmt"
)

func main() {
	stringStream := make(chan string)
	go func() {
		stringStream <- "Hello"
	}()

	fmt.Println(<-stringStream)
}
