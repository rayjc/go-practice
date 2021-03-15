package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	const numRoutines = 10
	const numData = 10

	// spawn 10 goroutines
	for i := 0; i < numRoutines; i++ {
		go send(c, numData)
	}

	// cannot range over unclosed channel
	// receive(c)

	for i := 0; i < numData*numRoutines; i++ {
		fmt.Println(i, " Received: ", <-c)
	}

	fmt.Println("about to exit")
}

func send(c chan<- int, n int) {
	for i := 0; i < n; i++ {
		c <- i
	}
}

func receive(c <-chan int) {
	count := 0
	for v := range c {
		fmt.Println(count, " Received: ", v)
		count++
	}
}
