package main

import "fmt"

func main() {
	var testValue = "Mhi"
	switch testValue {
	case "Daniel":
		fmt.Println("wassup daniel")
	case "Media":
		fmt.Println("wassup media")
	case "jenny":
		fmt.Println("wassup jenny")
	default:
		fmt.Println("default")
	}
}

/*
break isn't needed like in other languages

*/
