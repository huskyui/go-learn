package main

import "fmt"

func main() {
	i := 0
	for true {
		fmt.Println(i)
		if i > 10 {
			break
		}
		i++
	}
}
