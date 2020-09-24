package main

import "fmt"

func main() {
	greeting("adc", "apc")
}

func greeting(firstName string, secondName string) {
	fmt.Println(firstName, secondName)
}
