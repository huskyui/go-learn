package main

import "fmt"

func method1() {
	fmt.Println("1")
}

func method2() {
	fmt.Println("2")
}

func main() {
	defer method1()
	method2()
}
