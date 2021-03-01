package main

import "fmt"

func main() {
	done := make(chan bool)

	values := []string{"1", "2", "3"}

	for _, v := range values {
		u := v
		go func() {
			fmt.Println(u)
			done <- true
		}()
	}

	for _ = range values {
		<-done
	}
}
