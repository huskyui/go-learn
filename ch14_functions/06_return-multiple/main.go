package main

import "fmt"

func main() {
	var a1, a2 = exec(3, 2)
	fmt.Println(a1, a2)
}

func exec(num1, num2 int) (r1, r2 int) {
	r1 = num1 + num2
	r2 = num1 * num2
	return
}
