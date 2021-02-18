package main

import "fmt"

func greeting() func() string  {
	return func() string {
		return fmt.Sprintln("hello,top")
	}
}

func main() {
	greet := greeting();
	fmt.Println(greet())
	fmt.Printf("%T",greet)
}
