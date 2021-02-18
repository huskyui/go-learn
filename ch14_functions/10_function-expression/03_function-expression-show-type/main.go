package main

import "fmt"

func main() {
	greeting := func() {
		fmt.Println("hello,apc")
	}
	greeting()
	fmt.Printf("%T",greeting)
}
