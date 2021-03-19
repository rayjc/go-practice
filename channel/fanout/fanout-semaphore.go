
package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	q := make(chan bool)	// semaphore
	const numRoutines = 10
	const numData = 10

	// spawn 10 goroutines
	for i := 0; i < numRoutines; i++ {
		go send(c, numData, q)
	}
	
	go func() {
		// pull value from semaphore for however many goroutines
		for i := 0; i < numRoutines; i++ {
			<-q
		}
		close(c)
	}()
	
	receive(c)
	
	fmt.Println("about to exit")
}

func send(c chan<- int, n int, q chan<- bool) {
	for i := 0; i < n; i++ {
		c <- i
	}
	q <- true
}

func receive(c <-chan int) {
	count := 0
	for v := range c {
		fmt.Println(count, " Received: ", v)
		count++
	}
}
