package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func do_something(n uint64, c chan uint64) {
	defer wg.Done()
	fmt.Println("goes here with", n)
	time.Sleep(1)
	c <- n * 2
}

func main() {
	var array [10]uint64
	var c chan uint64 = make(chan uint64)
	for v := range array {
		wg.Add(1)
		go do_something(uint64(v), c)
		fmt.Printf("Hello world %d times!\n", v)
	}
	go func() {
		wg.Wait()
		close(c)
	}()

	for v := range c {
		fmt.Printf("receiving %d \n", v)
	}
}
