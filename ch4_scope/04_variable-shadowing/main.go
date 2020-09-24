package main

import "fmt"

/*
variable-shadowing

*/
func max(x int )int{
	return 42+x
}

func main()  {
	max := max(1)
	fmt.Println(max)
}