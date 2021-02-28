package main

import "fmt"

func main() {

	// 4,9 => 16,81
	for n := range sq(sq(gen(2, 3))) {
		fmt.Println(n)
	}

}

func gen(nums ...int) chan int {
	c := make(chan int)
	go func() {
		for _, num := range nums {
			c <- num
		}
		close(c)
	}()
	return c
}

func sq(ch chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range ch {
			out <- n * n
		}
		close(out)
	}()
	return out
}
