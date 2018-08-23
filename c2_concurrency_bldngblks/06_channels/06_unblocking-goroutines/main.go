package main

import (
	"fmt"
	"sync"
)

func main() {
	begin := make(chan interface{})
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			<-begin
			fmt.Printf("%v has begun\n", v)

		}(i)
	}
	fmt.Println("Unblocking go routines")
	close(begin)
	wg.Wait()
}
