package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	// concurrency
	go foo()
	go bar()

	wg.Wait()
}

func foo() {
	for i := 0; i < 45; i++ {
		fmt.Println("foo ", i)
		time.Sleep(3 * time.Millisecond)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 45; i++ {
		fmt.Println("bar ", i)
		time.Sleep(20 * time.Millisecond)
	}
	wg.Done()
}
