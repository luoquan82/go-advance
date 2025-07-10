package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg = sync.WaitGroup{}

func Add(c *int32) {
	defer wg.Done()
	for range 1000 {
		atomic.AddInt32(c, 1)
	}
}

func main() {
	var counter int32 = 0
	wg.Add(10)

	for range 10 {
		go Add(&counter)
	}

	wg.Wait()
	fmt.Println("Counter:", counter)
}
