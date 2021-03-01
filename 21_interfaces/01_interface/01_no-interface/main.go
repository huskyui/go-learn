package main

import "fmt"

type square struct {
	side float64
}

func (s square) area() float64 {
	return s.side * s.side
}

func main() {
	s := square{2.2}

	fmt.Println(s.area())
}
