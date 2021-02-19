package main

import "fmt"

type person struct {
	First string
	Last  string
	Age   int
}

type doubleZero struct {
	person
	First          string
	LicensseTokill bool
}

func (p person) Greeting() {
	fmt.Println("I'm just a regular person")
}

func (dz doubleZero) Greeting() {
	fmt.Println("Miss ,so good to see you")
}

func main() {
	p1 := doubleZero{
		person{
			"key", "value", 11,
		},
		"name",
		false,
	}

	p2 := person{
		"hey", "baby", 22,
	}

	// fmt.Println(p1.person.First)

	p1.Greeting()
	p2.Greeting()
}
