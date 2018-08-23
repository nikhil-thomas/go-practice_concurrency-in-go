package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan interface{})
	defer close(done)
	intStream := generator(done, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	pipeline := multiply(done, add(done, multiply(done, intStream, 10), 100), 10)

	for i := range pipeline {
		fmt.Println(i)
	}
}

func generator(done <-chan interface{}, integers ...int) <-chan int {
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for _, i := range integers {
			select {
			case <-done:
				return
			case intStream <- i:
				time.Sleep(100 * time.Millisecond)
			}
		}
		fmt.Println("generator done")
	}()
	return intStream
}

func multiply(done <-chan interface{}, intStream <-chan int, multiplier int) <-chan int {
	multipliedStream := make(chan int)
	go func() {
		defer close(multipliedStream)
		for i := range intStream {
			select {
			case <-done:
				return
			case multipliedStream <- i * multiplier:
			}
		}
		fmt.Println("multiply done")
	}()
	return multipliedStream
}

func add(done <-chan interface{}, intStream <-chan int, additive int) <-chan int {
	addedStream := make(chan int)
	go func() {
		defer close(addedStream)
		for i := range intStream {
			select {
			case <-done:
				return
			case addedStream <- i + additive:
			}
		}
		fmt.Println("add done")
	}()
	return addedStream
}
