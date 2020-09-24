package main

import "fmt"

func main() {
	greeting("as", "qs")
}

func greeting(fname, sname string) {
	fmt.Println(fname, sname)
}
