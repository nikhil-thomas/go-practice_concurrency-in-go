package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := sync.NewCond(&sync.Mutex{})
	queue := make([]interface{}, 0, 0)

	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()
		fmt.Println(queue[0])
		queue = queue[1:]
		fmt.Println("Removed from queue")
		c.L.Unlock()
		c.Signal()
	}
	for i := 0; i < 10; i++ {
		c.L.Lock()
		for len(queue) == 3 {
			fmt.Println("wait")
			c.Wait()
		}
		fmt.Println("add", i)
		queue = append(queue, i)
		go removeFromQueue(1 * time.Second)
		c.L.Unlock()
	}
}
