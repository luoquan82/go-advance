package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func sender(ch chan<- int) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		ch <- i + 1
		fmt.Println("Sent:", i+1)
	}
	close(ch)
}

func receiver(ch <-chan int) {
	defer wg.Done()
	for num := range ch {
		fmt.Println("Received:", num)
	}
}

func main() {
	ch := make(chan int)

	wg.Add(2)
	go sender(ch)
	go receiver(ch)

	wg.Wait()
}
