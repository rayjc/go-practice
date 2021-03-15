package main

import (
	"fmt"
)

func main() {
	// create channel (blocking queue)
	c := make(chan int)

	go func() {
		c <- 42
		// blocks until channel is empty
	}()

	fmt.Println(<-c)
}
