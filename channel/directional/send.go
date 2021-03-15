package main

import (
	"fmt"
)

func main() {
	// create  channcel
	c := make(chan int)

	// convert to a send/push only channel
	go func(cs chan<- int) {
		cs <- 42
		fmt.Printf("------\n")
		fmt.Printf("cs\t%T\n", cs)
	}(c)

	fmt.Println(<-c)

	fmt.Printf("------\n")
	fmt.Printf("c\t%T\n", c)
}
