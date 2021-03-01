package main

import "fmt"

func main() {
	age := 24
	fmt.Println(&age)
	changeMe(&age)
	fmt.Println(age)
}

func changeMe(age *int) {
	fmt.Println(age)
	*age = 48

}
