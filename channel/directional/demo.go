package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go send(c)

	receive(c)

	fmt.Println("about to exit")
}

func send(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	// close send channel
	close(c)
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println("Received: ", v)
	}
}

// func main() {
// 	c := gen()
// 	receive(c)

// 	fmt.Println("about to exit")
// }

// func gen() <-chan int {
// 	c := make(chan int)

// 	go func() {
// 		for i := 0; i < 100; i++ {
// 			c <- i
// 		}
// 		// close send channel
// 		close(c)
// 	}()

// 	return c
// }

// func receive(c <-chan int) {
// 	for v := range c {
// 		fmt.Println("Received: ", v)
// 	}
// }