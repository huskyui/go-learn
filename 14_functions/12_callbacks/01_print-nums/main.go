package main

import "fmt"

func visit(numbers []int, callback func(int)) {
	for _, v := range numbers {
		callback(v)
	}
}

// 回调方法，有点像java里面的匿名内部类
func main() {
	numbers := []int{1, 2, 3, 5, 6}
	visit(numbers, func(i int) {
		fmt.Println(i)
	})
}
