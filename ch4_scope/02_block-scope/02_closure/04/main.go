package main
import "fmt"

func wrapper() func() int {
	x:=0
	fmt.Println("wrapper:",x)
	return func() int {
		fmt.Println("func: ",x)
		x++
		return x
	}
}
func main() {
	increment:=wrapper()
	increment1:=wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment1())
	fmt.Println(increment1())
}
