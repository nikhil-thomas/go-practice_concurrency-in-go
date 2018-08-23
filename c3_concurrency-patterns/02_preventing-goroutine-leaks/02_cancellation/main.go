package main

import (
	"fmt"
	"time"
)

func main() {
	doWork := func(done <-chan interface{}, strings <-chan string) <-chan interface{} {
		terminated := make(chan interface{})
		go func() {
			defer fmt.Println("dowork exited")
			defer close(terminated)
			for {
				select {
				case <-done:
					return
				case s := <-strings:
					fmt.Println(s)
				}
			}
		}()
		return terminated
	}
	done := make(chan interface{})
	terminated := doWork(done, nil)

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("cancelling dowork roroutine")
		close(done)
	}()

	<-terminated

	fmt.Println("Done")
}
