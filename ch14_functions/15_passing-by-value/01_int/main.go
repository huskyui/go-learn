package main

import "fmt"

func main() {
	age := 24
	changeAge(age)
	fmt.Println(age)
	
}

func changeAge(x int){
	x = 48
}
