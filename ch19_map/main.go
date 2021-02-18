package main

import "fmt"

func main() {
	// not init map
	// var myGreeting map[string]string
	// fmt.Println(myGreeting)
	// fmt.Println(myGreeting == nil)

	// myGreeting["tim"] = "Good morning"
	// myGreeting["Jenny"] = "Bon"

	// fmt.Println(myGreeting)

	// myGreeting := make(map[string]string)
	// myGreeting["tim"] = "good"
	// myGreeting["Jenny"] = "bon"
	// fmt.Println(myGreeting)
	// fmt.Println(myGreeting["tim"])
	// fmt.Println(myGreeting["add"])

	// myGreeting := map[string]string{}
	// myGreeting["tim"] = "Good morning"
	// myGreeting["jenny"] = "Bon"
	// fmt.Println(myGreeting)

	// myGreeting := map[string]string{
	// 	"Tim":   "Good morning",
	// 	"Jenny": "Bonjour",
	// }

	// fmt.Println(myGreeting["Jenny"])

	// myGreeting := map[string]string{
	// 	"Tim":   "Good morning",
	// 	"jenny": "bonjour",
	// }

	// myGreeting["Harleen"] = "Howddy"
	// fmt.Println(myGreeting)
	// // fmt.Println(len(myGreeting))
	// myGreeting["Harleen"] = "Giddy"
	// fmt.Println(myGreeting)

	// delete(myGreeting, "Harleen")
	// fmt.Println(myGreeting)

	// myGreeting := map[int]string{
	// 	0: "0",
	// 	1: "1",
	// 	2: "2",
	// 	3: "3",
	// }
	// fmt.Println(myGreeting)

	// if val, exists := myGreeting[2]; exists {
	// 	fmt.Println("The value exists")
	// 	fmt.Println("val: ", val)
	// 	fmt.Println("exists: ", exists)
	// } else {
	// 	fmt.Println("That value doesn't exist.")
	// 	fmt.Println("val: ", val)
	// 	fmt.Println("exists: ", exists)
	// }
	// fmt.Println(myGreeting)

	// if val, exists := myGreeting[7]; exists {
	// 	delete(myGreeting, 7)
	// 	fmt.Println("val", val)
	// 	fmt.Println("exists", exists)
	// } else {
	// 	fmt.Println("the value doesn't exists")
	// 	fmt.Println("val: ", val)
	// 	fmt.Println("eixsts: ", exists)
	// }
	// fmt.Println(myGreeting)

	// myGreeting := map[int]string{
	// 	1: "Gooding",
	// 	2: "2234",
	// 	3: "2ass23",
	// }

	// for key, val := range myGreeting {
	// 	fmt.Println(key, val)

	// }

	// letter := 'A'
	// fmt.Println(letter)
	// fmt.Printf("%T \n", letter)

	// fmt.Println(rune("A"[0]))
	// fmt.Println(rune("Hello"[0]))

	// n := hashBucket("Go", 12)
	// fmt.Println(n)

	buckets := make([]int, 1)
	fmt.Println(buckets[0])
	buckets[0] = 42
	fmt.Println(buckets[0])
	buckets[0]++
	fmt.Println(buckets[0])
	buckets[1] = 2
}

// func hashBucket(word string, buckets int) int {
// 	letter := int(word[0])
// 	bucket := letter % buckets
// 	return bucket
// }
