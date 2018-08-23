package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.NewTicker(1 * time.Second)
	i := 0
	for v := range t.C {
		fmt.Println(v)
		if i++; i >= 10 {
			// t.Stop()
			break
		}
	}
	fmt.Println("end", <-t.C)
}
