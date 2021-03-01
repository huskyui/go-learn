package main

import "fmt"

func main() {
	mySlice := []string{"Monday", "Tuesday"}
	myOtherSlice := []string{"Wenednesday", "Thurday", "Friday"}

	mySlice = append(mySlice, myOtherSlice...)
	fmt.Println(mySlice)
	fmt.Println(mySlice[:2])
	mySlice = append(mySlice[:2], myOtherSlice[3:]...)
	fmt.Println(mySlice)

}
