package main

import "fmt"

func main() {
	var numOne int
	var numTwo int
	fmt.Scan(&numOne)
	fmt.Scan(&numTwo)
	fmt.Println(numOne, "%", numTwo, "=", numOne%numTwo)
}
