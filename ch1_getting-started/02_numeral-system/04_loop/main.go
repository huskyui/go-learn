package main

import "fmt"

func main() {
	for i := 10000; i < 10100; i++ {
		fmt.Printf("%d\t%b\t%x\t\n", i, i, i)
	}
}
