package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func printNumbers(threadName string, start, end int) {
	for i := start; i <= end; i++ {
		fmt.Printf("%s print %d\n", threadName, i)
		time.Sleep(100 * time.Millisecond)
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go printNumbers("Thread1", 1, 10)
	go printNumbers("Thread2", 2, 10)
	wg.Wait()
}
