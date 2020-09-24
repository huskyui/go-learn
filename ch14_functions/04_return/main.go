package main

import "fmt"

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(fullName("ad", "ap"))
}

func add(num1, num2 int) int {
	return num1 + num2
}

func fullName(name1, name2 string) string {
	return fmt.Sprint(name1, name2)
}
