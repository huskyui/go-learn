package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var count int

func main() {

	wg.Add(2)
	go incrementor("bar")
	go incrementor("foo")
	wg.Wait()

}

func incrementor(s string) {
	// create random number
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 20; i++ {
		x := count
		x++
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		count = x
		fmt.Println(s, i, "counter", count)
	}

	wg.Done()

}
