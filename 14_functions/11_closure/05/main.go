package main

import "fmt"

func wrapper() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

func main() {
	incrementA := wrapper()
	incrementB := wrapper()
	fmt.Println(incrementA())
	fmt.Println(incrementA())
	fmt.Println(incrementB())
	fmt.Println(incrementB())
	fmt.Println(incrementB())
}
