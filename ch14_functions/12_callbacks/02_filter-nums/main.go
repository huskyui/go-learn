package main

import "fmt"

func filterNums(numbers []int,callback func(int) bool) []int {
	var nums []int;
	for _,v :=range numbers{
		if callback(v){
			nums = append(nums,v)
		}
	}
	return nums
}

func main() {
	nums := []int{1,2,3,4,5,6,7}

	fmt.Println(filterNums(nums, func(i int) bool {
		if i%2 == 0 {
			return true
		} else {
			return false
		}
	}))
}
