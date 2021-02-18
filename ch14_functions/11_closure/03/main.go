package main

import "fmt"

func main() {
	var x int
	increment := func() {
		x++
		fmt.Println(x)
	}

	increment()
	increment()
}
