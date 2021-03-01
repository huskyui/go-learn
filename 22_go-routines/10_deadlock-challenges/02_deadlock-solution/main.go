package main

import "fmt"

func main() {
	c := make(chan int)
	// make new goroutine that write msg to channel
	go func() {
		c <- 1
	}()
	fmt.Println(<-c)
}
