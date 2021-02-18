package main

import "fmt"

func main() {
	foo(1, 2)
	foo(1)
	aSlice := []int{1, 2, 4}
	foo(aSlice...)
	foo()
}

func foo(numbers ...int) {
	fmt.Println(numbers)
}
