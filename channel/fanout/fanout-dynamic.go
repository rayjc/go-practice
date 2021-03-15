package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	input := make(chan int)
	transformed := make(chan int)

	go send(input, 100)

	// Note: this must be a goroutine so receive can be reached and consume data
	go fanout(input, transformed)

	// blocking until transformed closes
	receive(transformed)

	fmt.Println("about to exit")
}

func send(c chan<- int, n int) {
	for i := 0; i < n; i++ {
		c <- i
	}
	close(c)
}

func fanout(c1 <-chan int, c2 chan<- int) {
	var wg sync.WaitGroup
	maxGo := 0

	for v := range c1 {
		wg.Add(1)
		go func(v int) {
			// do some work from each value in pipeline
			c2 <- v << 1

			numGo := runtime.NumGoroutine()
			if maxGo < numGo {
				maxGo = numGo
			}
			wg.Done()
		}(v)
	}

	wg.Wait()
	fmt.Println(maxGo)
	close(c2)
}

func receive(c <-chan int) {
	count := 0
	for v := range c {
		fmt.Println(count, " Received: ", v)
		count++
	}
}
