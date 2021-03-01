package main

import "fmt"

func main() {
	name := "huskyui"
	fmt.Println(&name)
	changeName(&name)
	fmt.Println(name)
}

func changeName(name *string) {
	fmt.Println(name)
	fmt.Println(*name)
	*name = "adios"
}
