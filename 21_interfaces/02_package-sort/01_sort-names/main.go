package main

import (
	"fmt"
	"sort"
)

// type people []string

// func (p people) Len() int {
// 	return len(p)
// }

// func (p people) Less(i, j int) bool {
// 	return p[i] < p[j]
// }

// func (p people) Swap(i, j int) {
// 	p[i], p[j] = p[j], p[i]
// }

// func main() {
// 	studyGroup := people{"b", "a", "c"}
// 	sort.Sort(studyGroup)
// 	fmt.Println(studyGroup)
// }

// func main() {
// 	s := []string{"b", "a", "c"}
// 	sort.Sort(sort.StringSlice(s))
// 	fmt.Println(s)
// }

// func main() {
// 	// s := []string{"b", "c", "c"}
// 	// sort.Strings(s)
// 	// fmt.Println(s)

// 	// type a int;
// }

type person struct {
	Name string
	Age  int
}

func (p person) toString() string {
	return p.Name + string(p.Age)
}

type ByAge []person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
	people := []person{
		{"Bob", 18},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}

	fmt.Println(people[0])

	sort.Sort(ByAge(people))
	fmt.Println(people)
}
