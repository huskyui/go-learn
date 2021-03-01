package main

import (
	"encoding/json"
	"fmt"
)

// type person struct {
// 	First       string
// 	Last        string
// 	Age         int
// 	notExported int
// }

// type person struct {
// 	first string
// 	last  string
// 	age   int
// }

// type person struct {
// 	First string
// 	Last  string `json:"-"`
// 	Age   int    `json:"wisdom score"`
// }

// func main() {
// 	// p1 := person{"first", "last", 20, 1}
// 	// bs, _ := json.Marshal(p1)
// 	// fmt.Println(bs)
// 	// fmt.Printf("%T \n", bs)
// 	// fmt.Println(string(bs))

// 	// p1 := person{"james", "bond", 20}

// 	// fmt.Println(p1)
// 	// bs, _ := json.Marshal(p1)
// 	// fmt.Println(string(bs))

// 	p1 := person{"James", "Bond", 20}
// 	fmt.Println(p1)
// 	bs, _ := json.Marshal(p1)
// 	fmt.Println(string(bs))
// }

// type person struct {
// 	First string
// 	Last  string
// 	Age   int
// }

// func main() {
// 	var p1 person
// 	fmt.Println(p1.First)
// 	fmt.Println(p1.Last)
// 	fmt.Println(p1.Age)

// 	bs := []byte(`{"First":"James","Last":"Bond","Age":20}`)
// 	json.Unmarshal(bs, &p1)

// 	fmt.Println(bs)
// 	fmt.Println(string(bs))

// 	fmt.Println("------")
// 	fmt.Println(p1.First)
// 	fmt.Println(p1.Last)
// 	fmt.Println(p1.Age)
// 	fmt.Printf("%T \n", p1)
// }

type person struct {
	First string
	Last  string
	Age   int `json:"wisdom score"`
}

func main() {
	var p1 person

	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)

	bs := []byte(`{"First":"Json","Last":"Bond","wisdom score":20}`)

	json.Unmarshal(bs, &p1)

	fmt.Println(p1)

}
