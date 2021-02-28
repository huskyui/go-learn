package main

func main() {

	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

	//fmt.Println(<-c)

	//for n:=range c{
	//	fmt.Println(n)
	//}
	//for {
	//	fmt.Println(<-c)
	//}

}
