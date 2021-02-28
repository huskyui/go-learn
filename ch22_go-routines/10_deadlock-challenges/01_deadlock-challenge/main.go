package main

import "fmt"

func main() {
	// 同一个goroutine中使用同一个channel读写
	c := make(chan int)
	c <- -1
	fmt.Println(<-c)
}
