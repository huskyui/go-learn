package main

import "fmt"

func main() {
	variable := "123"
	switch variable {
	case "123", "234":
		fmt.Println("here is 123,234")
	case "345", "124", "1024":
		fmt.Println("here is 1024")
	}
}
