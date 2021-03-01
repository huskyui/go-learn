package main

import "fmt"

func main() {

	done := make(chan bool)
	values := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	// v好像是共享的
	for _, v := range values {
		go func() {
			fmt.Println(v)
			done <- true
		}()
	}

	for _ = range values {
		<-done
	}

}
