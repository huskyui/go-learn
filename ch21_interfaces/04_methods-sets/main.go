package main

import (
	"fmt"
	"math"
)

// type circle struct {
// 	radius float64
// }

// // value-receiver_value-type
// type shape interface {
// 	area() float64
// }

// func (c circle) area() float64 {
// 	return math.Pi * c.radius * c.radius
// }

// type shapeV1 interface {
// 	areaV1() float64
// }

// func (c *circle) areaV1() float64 {
// 	return math.Pi * c.radius * c.radius
// }

// func infoV1(shape shapeV1) {
// 	fmt.Println(shape.areaV1())
// }

// func info(s shape) {
// 	fmt.Println("area", s.area())
// }

// func main() {
// 	c := circle{5}
// 	info(c)
// 	infoV1(&c)
// }

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (c *circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func info(s shape) {
	fmt.Println("area", s.area())
}

func main() {
	c := circle{5}
	info(&c)
}
