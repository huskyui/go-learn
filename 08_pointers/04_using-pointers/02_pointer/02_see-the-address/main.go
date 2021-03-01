package main

import "fmt"

func zero(z *int) {
	// z is the pointer
	fmt.Println(z)
	*z = 0
}

func main() {
	x := 5
	fmt.Println(&x)
	zero(&x)
	fmt.Println(x)
}
