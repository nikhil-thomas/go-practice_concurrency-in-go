package main

import (
	"fmt"
)

func main() {
	done := make(chan interface{})
	defer close(done)
	var message string

	for token := range toString(done, take(done, repeat(done, "I", "AM", "E"), 10)) {
		message += (token + " ")
	}
	fmt.Printf("message %s\n", message)
}

func repeat(done <-chan interface{}, values ...interface{}) <-chan interface{} {
	valueStream := make(chan interface{})
	go func() {
		defer close(valueStream)
		for {
			for _, v := range values {
				select {
				case <-done:
					return
				case valueStream <- v:
				}
			}
		}

	}()
	return valueStream
}

func take(done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{} {
	takeStream := make(chan interface{})
	go func() {
		defer close(takeStream)
		for i := 0; i < num; i++ {
			select {
			case <-done:
				return
			case v := <-valueStream:
				takeStream <- v
			}
		}
	}()
	return takeStream
}

func toString(done <-chan interface{}, valueStream <-chan interface{}) <-chan string {
	stringStream := make(chan string)
	go func() {
		defer close(stringStream)
		for value := range valueStream {
			select {
			case <-done:
				return
			case stringStream <- value.(string):
			}
		}
	}()
	return stringStream
}
