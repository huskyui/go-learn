package main

import (
	"fmt"
	"strconv"
)

func main() {
	// // int to float
	// // var x = 12
	// var y = 12.123
	// fmt.Println(int(y))
	// // fmt.Println(y + float64(x))

	// rune to string
	// var x rune = 'a'
	// fmt.Println(string(x))

	// var y byte = 'b'
	// fmt.Println(string(y))

	// s := "community"
	// s := "中文"

	// byteArr := []byte(s)
	// // fmt.Println([]byte(s))
	// // fmt.Println([]rune(s))

	// fmt.Println(byteArr[:3])

	// fmt.Println(s[:3])
	// var x = "12_1"
	// z, err := strconv.Atoi(x)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(z)

	// x := 12

	// y := "i have this many " + strconv.Itoa(x)
	// fmt.Println(y)

	b, _ := strconv.ParseBool("true")
	fmt.Println(b)
	f, _ := strconv.ParseFloat("3.1415", 32)
	fmt.Println(f)
}
