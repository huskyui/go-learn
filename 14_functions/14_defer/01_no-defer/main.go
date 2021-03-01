package main

import "fmt"

func method1() {
	fmt.Println("1")
}

func method2() {
	fmt.Println("2")
}

func method3() {
	fmt.Println("3")
}

func main() {
	method1()
	method3()
	method2()
}
