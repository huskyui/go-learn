package main

import "fmt"

var x int

func increment()  {
	x++
	fmt.Println(x)
}

func main() {
	increment()
	increment()
	increment()
}
