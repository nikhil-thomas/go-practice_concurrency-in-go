package main

import (
	"fmt"
	"sync"
)

func main() {
	type Button struct {
		Clicked *sync.Cond
	}

	button := Button{
		Clicked: sync.NewCond(&sync.Mutex{}),
	}

	subscribe := func(c *sync.Cond, fn func()) {
		var goroutineRunning sync.WaitGroup
		goroutineRunning.Add(1)
		go func() {
			goroutineRunning.Done()
			c.L.Lock()
			defer ec.L.Unlock()
			c.Wait()
			fn()
		}()
		goroutineRunning.Wait()
	}

	var clickRegisterd sync.WaitGroup
	clickRegisterd.Add(3)
	subscribe(button.Clicked, func() {
		fmt.Println("Maximizing window")
		clickRegisterd.Done()
	})
	subscribe(button.Clicked, func() {
		fmt.Println("Displaying annoying dialogue box")
		clickRegisterd.Done()
	})
	subscribe(button.Clicked, func() {
		fmt.Println("Mouse Clicked")
		clickRegisterd.Done()
	})
	button.Clicked.Broadcast()
	clickRegisterd.Wait()

}
