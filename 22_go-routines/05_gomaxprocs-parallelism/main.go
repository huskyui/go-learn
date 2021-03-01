package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	numCpu := runtime.NumCPU()
	fmt.Println("current process  the number of logical cpus", numCpu)
	runtime.GOMAXPROCS(numCpu)
}

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()

}

func foo() {
	for i := 0; i < 45; i++ {
		fmt.Println("foo", i)
		time.Sleep(2 * time.Millisecond)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 45; i++ {
		fmt.Println("bar", i)
		time.Sleep(20 * time.Millisecond)
	}
	wg.Done()
}
