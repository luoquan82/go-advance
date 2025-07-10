package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var lock = sync.Mutex{}

func Add(c *int) {
	lock.Lock()
	defer lock.Unlock()
	for range 1000 {
		*c++
	}
	wg.Done()
}

func main() {
	var counter = 0
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go Add(&counter)
	}
	wg.Wait()
	fmt.Println("Counter:", counter)
}
