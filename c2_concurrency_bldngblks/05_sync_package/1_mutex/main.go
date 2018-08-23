package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	var lock sync.Mutex

	increment := func() {
		lock.Lock()
		defer lock.Unlock()
		count++
		fmt.Printf("Incrementing: %d\n", count)
	}

	decrement := func() {
		lock.Lock()
		defer lock.Unlock()
		count--
		fmt.Printf("Incrementing: %d\n", count)
	}

	var arithematic sync.WaitGroup
	for i := 0; i <= 5; i++ {
		arithematic.Add(1)
		go func() {
			defer arithematic.Done()
			increment()
		}()
	}

	for i := 0; i <= 5; i++ {
		arithematic.Add(1)
		go func() {
			defer arithematic.Done()
			decrement()
		}()
	}

	arithematic.Wait()
	fmt.Println("Arithematic Done")
}
