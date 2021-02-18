package main

import "fmt"

func main() {
	//羌笛吹落梅让人分不清异乡和故里
	var X [256]byte
	fmt.Println(len(X))
	fmt.Println(X[42])
	for i := 0; i < 256; i++ {
		X[i] = byte(i)
	}
	for i, v := range X {
		fmt.Printf("%v-%T-%b\n", v, v, v)
		if i > 50 {
			break
		}
	}

}
