package main

import (
	"fmt"
)

func main() {
	// semaphore
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		x := <-done
		y := <-done

		fmt.Println("x ", x)
		fmt.Println("y ", y)
		close(c)

	}()

	for n := range c {
		fmt.Println(n)
	}

}
