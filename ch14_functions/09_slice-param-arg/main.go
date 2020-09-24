package main

import "fmt"

func main() {
	data := []float64{47, 23, 99, 12}
	fmt.Println(average(data))

}

func average(data []float64) float64 {
	total := 0.0
	for _, v := range data {
		total += v
	}
	return total / float64(len(data))
}
