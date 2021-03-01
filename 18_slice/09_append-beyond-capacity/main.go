package main

import "fmt"

func main() {
	greeting := make([]string, 3, 5)
	greeting[0] = "0"
	greeting[1] = "1"
	greeting[2] = "2"
	greeting = append(greeting, "3")
	greeting = append(greeting, "4")
	greeting = append(greeting, "5")
	greeting = append(greeting, "6")

	fmt.Println(greeting[6])
	fmt.Println(len(greeting))
	fmt.Println(cap(greeting))
}
