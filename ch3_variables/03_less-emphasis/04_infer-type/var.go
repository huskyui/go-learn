package main

import "fmt"

func main() {
	var message = "hello,world"
	var a, b, c, d = 1, 2, 3, "abc"
	fmt.Println(message, a, b, c, d)
	fmt.Printf("%T %T",a,d)
}
