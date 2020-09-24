package main

import "fmt"

func main() {
	n := average(43, 56, 87, 12)
	fmt.Println(n)
}

func average(sf ...float64) float64 {
	fmt.Println(sf)
	fmt.Printf("%T \n", sf)
	var total float64
	for i := 0; i < len(sf); i++ {
		total += sf[i]
	}
	return total / float64(len(sf))
}
