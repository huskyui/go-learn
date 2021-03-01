package main

import "fmt"

func main() {
	if false {
		fmt.Println("first statement")
	} else if true {
		fmt.Println("second statement")
	} else {
		fmt.Println("third statement")
	}
}
