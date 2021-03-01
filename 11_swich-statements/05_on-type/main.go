package main

import "fmt"

type concat struct {
	greeting string
	name     string
}

func SwitchOnType(a interface{}) {
	switch a.(type) {
	case int:
		fmt.Println("int", a)
	case float64:
		fmt.Println("float64", a)
	case concat:
		fmt.Println("concat", a)
	}
}

func main() {
	var a int = 1
	var b float64 = 1.2
	c := concat{
		greeting: "你好",
		name:     "huskyui",
	}
	SwitchOnType(a)
	SwitchOnType(b)
	SwitchOnType(c)

}
