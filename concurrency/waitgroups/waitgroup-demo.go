package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		fmt.Println("In first goroutine")
		wg.Done()
	}()

	go func() {
		fmt.Println("In second goroutine")
		wg.Done()
	}()

	fmt.Println("end of main")

	wg.Wait()
}
