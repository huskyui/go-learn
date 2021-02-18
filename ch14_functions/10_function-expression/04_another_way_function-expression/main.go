package main

import "fmt"

func getFunc() func() string{
	return func() string{
		return fmt.Sprintln("hello,jungle")
	}
}

func main() {
	greeting := getFunc()
	fmt.Println(greeting())
}
