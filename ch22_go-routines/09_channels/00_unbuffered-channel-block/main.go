package main

import (
	"fmt"
	"time"
)

func main() {

	// chan channel
	// ch <- v 将数据v发送给channel ch中
	// c := <- ch 从channel ch读取数据，并赋值c
	// c,ok := <- ch
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- (i + 1)
		}
	}()

	go func() {
		for {
			reciveData, ok := <-c
			fmt.Println("channel is open", ok, "receive data", reciveData)
		}
	}()

	time.Sleep(time.Second)

}
