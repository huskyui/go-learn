package main

import (
	"fmt"
)

type customer struct {
	name string
	age int
}

func main() {
	c1 := customer{"todd",11}
	fmt.Println(c1.name)
	changeMe2(c1)
	fmt.Println(c1)

}

func changeMe2(z customer)  {
	fmt.Println(&z)
	z.age = 12



}


//func changeMe(z *customer){
//	fmt.Println(z)
//	fmt.Println(&z.name)
//	fmt.Println(&z.age)
//	z.name = "Rocky"
//	fmt.Println(z)
//	fmt.Println(&z.name)
//}
