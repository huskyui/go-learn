package main

import "fmt"

func main() {
	var s = greet("1", "2")
	fmt.Println(s)
}
func greet(name1, name2 string) (s string) {
	s = fmt.Sprintln(name1, name2)
	return
}
