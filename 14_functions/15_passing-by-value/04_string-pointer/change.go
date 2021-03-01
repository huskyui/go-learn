package main

import "fmt"

func main() {
	name := "Todd"
	fmt.Println(&name) // 0xc0000441f0
	fmt.Print(*&name)
	changeMe(&name)
	fmt.Println(name)

}

func changeMe(z *string) {
	fmt.Println(z)
	fmt.Println(*z)
	*z = "Rocky"
	fmt.Println(z)
	fmt.Println(*z)
}
