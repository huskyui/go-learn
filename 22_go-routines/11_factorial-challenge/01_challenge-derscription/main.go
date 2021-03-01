package main

import "fmt"

//func main() {
//	f:=factorial(4)
//	fmt.Println(f)
//}
//
//
//func factorial(n int) int{
//	total := 1
//	for i:=n;i>0;i--{
//		total *=i
//	}
//	return total
//}

func main() {

	// 在使用goroutine+channel来实现方法
	c := factorial(4)
	for n := range c {
		fmt.Println(n)
	}

}

func factorial(n int) chan int {
	c := make(chan int)

	go func() {
		total := 1
		for i := n; i > 0; i-- {
			total *= i
		}
		c <- total
		close(c)
	}()
	return c
}
