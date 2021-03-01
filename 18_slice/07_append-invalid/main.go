package main

import "fmt"

func main() {
	greeting := make([]string, 3, 5)

	greeting[0] = "Gooding mo"
	greeting[1] = "a"
	greeting[2] = "b"
	greeting[3] = "c"
	//invalid
	// greeting = append(greeting, "3")
	fmt.Println(greeting[3])
}
