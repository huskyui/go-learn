package main

import "fmt"

func main() {
	for i := 48; i < 149; i++ {
		fmt.Printf("%v -%v -%v \n", i, string(i), []byte(string(i)))
	}
}
