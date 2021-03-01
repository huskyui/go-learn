package main

import "fmt"

func main() {

	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}

	// remember close your channel
	// if you do not close channel,you will receive this error
	// fatal error: all goroutines are asleep - deadlock!

	// you need go version 1.5.2 or greater?????
	// otherwise you will receive this error??????

}
