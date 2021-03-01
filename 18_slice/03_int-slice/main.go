package main

import "fmt"

func main() {
	mySlice := make([]int, 0, 3)
	fmt.Println("-------------")
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	// capacity will auto increase.

	for i := 0; i < 80; i++ {
		mySlice = append(mySlice, i)
		fmt.Println("Len: ", len(mySlice), "Capacity ", cap(mySlice), "Value: ", mySlice[i])
	}

}
