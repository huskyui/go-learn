package main

import (
	"fmt"
	"testing"
)

func BenchmarkHell1(b *testing.B) {
	count := 0
	for i := 0; i < b.N; i++ {
		if i%2 == 0 {
			count += i
		} else if i%5 == 0 {
			count += i
		}
	}
	fmt.Println(count)
}
