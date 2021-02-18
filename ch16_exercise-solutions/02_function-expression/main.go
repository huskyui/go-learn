package main

import "fmt"

func main() {
	// 创建一个方法，并且使用
	half := func(n int) (int, bool) {
		return n / 2, n%2 == 0
	}
	fmt.Println(half(5))
}
