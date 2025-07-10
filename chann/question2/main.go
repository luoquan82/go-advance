package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func send(ch chan<- int) {
	for i := range 100 {
		ch <- i + 1
		fmt.Printf("Send %d to channel\n", i+1)
	}
	close(ch)
}

func receive(ch <-chan int) {
	for d := range ch {
		fmt.Println("Received:", d)
	}
}

func main() {
	ch := make(chan int, 100)

	wg.Add(2)
	go func() {
		defer wg.Done()
		send(ch)
	}()

	go func() {
		defer wg.Done()
		receive(ch)
	}()

	wg.Wait()
}
