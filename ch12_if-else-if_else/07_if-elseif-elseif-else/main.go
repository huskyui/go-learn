package main

import "fmt"

func main() {
	if false {
		fmt.Println("1")
	} else if false {
		fmt.Println("2")
	} else if true {
		fmt.Println(3)
	} else {
		fmt.Println(4)
	}
}
