package main

import "fmt"

func main() {
	mySlice := []string{"Monday", "Tuesday"}
	myOtherSlice := []string{"Wenednesday", "Thurday", "Friday"}
	mySlice = append(mySlice, myOtherSlice...)
	fmt.Println(mySlice)

}
