package main

import "fmt"

func main() {
	greeting("huskyui")
	greeting("adios")
}

func greeting(name string) {
	fmt.Println("hello ", name)
}
