package main

import "fmt"

func main() {
	name := "huskyui"
	changeName(name)
	fmt.Println(name)
}

func changeName(name string){
	name = "adios"
}
