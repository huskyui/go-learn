package main

import (
	"fmt"
	"math/rand"
	"sync"
)

const numFactorials = 100
const rdLimit = 20

func main() {
	var w sync.WaitGroup
	// use countDownLatch solution
	w.Add(numFactorials)
	factorial(&w)
	w.Wait()
}

func factorial(wmain *sync.WaitGroup) {
	var w sync.WaitGroup
	rand.Seed(42)

	w.Add(numFactorials + 1)

	for j := 1; j <= numFactorials; j++ {
		// create 100 routines
		go func() {
			f := rand.Intn(rdLimit)
			w.Wait()
			total := 1

			for i := f; i > 0; i-- {
				total *= i
			}
			fmt.Println(f, total)
			(*wmain).Done()
			// todo use channel
		}()
		w.Done()
	}
	fmt.Println("all done with initialization")
	w.Done()

}
