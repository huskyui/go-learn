package main

import "fmt"

func main() {
	myFirstName := "adios"
	switch  {
	case len(myFirstName)==5:
		fmt.Println("congratulation")
	case myFirstName == "adios":
		fmt.Println("congratulation1")
	case myFirstName == "aaa":
		fmt.Println("congratulation2")



	}

}
