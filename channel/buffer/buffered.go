package main

import (
	"fmt"
)

func main() {
	// create channel (blocking queue)
	//
	c := make(chan int, 1)

	fmt.Println(cap(c))
	fmt.Println(len(c))

	c <- 42
	// blocks only when size/len reaches allocated capacity

	fmt.Println(cap(c))
	fmt.Println(len(c))

	fmt.Println(<-c)
}
