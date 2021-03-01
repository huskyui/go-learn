package main

import "fmt"

func main() {
	// var name interface{} = "Sydney"
	// str, ok := name.(string)
	// if ok {
	// 	fmt.Println("%q\n", str)
	// } else {
	// 	fmt.Printf("value is not a string \n")
	// }
	var x interface{} = 7
	// fmt.Println(x.(int) + 6)
	// fmt.Println(int(x))
	fmt.Println(x.(int))
}
