package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func taskScheduler(tasks []func()) {
	for i, task := range tasks {
		go func(idx int, t func()) {
			defer wg.Done()
			fmt.Printf("Starting task %d...\n", idx+1)
			start := time.Now()
			t()
			fmt.Printf("Task %d executed in %s\n", idx+1, time.Since(start))
		}(i, task)
	}
}

func main() {
	tasks := []func(){
		func() {
			r := rand.Intn(10)
			time.Sleep(time.Millisecond * time.Duration(r*1000))
		},
		func() {
			r := rand.Intn(10)
			time.Sleep(time.Millisecond * time.Duration(r*1000))
		},
		func() {
			r := rand.Intn(10)
			time.Sleep(time.Millisecond * time.Duration(r*1000))
		},
	}

	wg.Add(len(tasks))
	taskScheduler(tasks)
	wg.Wait()
}
