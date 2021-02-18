package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["a"] = 123
	m["b"] = 245
	fmt.Println(m)
	changeMe(m)
	fmt.Println(m)
}

func changeMe(z map[string]int){
	z["todd"] = 44
}
