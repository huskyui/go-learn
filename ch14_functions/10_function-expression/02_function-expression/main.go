package main

import "fmt"

func main() {
	greeting := func() {
		fmt.Println("hello,apc")
	}
	fmt.Printf("%T",greeting)
	greeting()
}
