package main

import "fmt"

type person struct {
	First string
	Last  string
	Age   int
}

type doubleZero struct {
	person
	LicenseToKill bool
}

func main() {
	p1 := doubleZero{
		person: person{
			First: "JAMES",
			Last:  "bOND",
			Age:   1,
		},
		LicenseToKill: true,
	}

	p2 := doubleZero{
		person: person{
			"james",
			"bond",
			1,
		},
		LicenseToKill: false,
	}

	fmt.Println(p1)
	fmt.Println(p2)

}
