package main

import "fmt"

func main() {
	var x [5]int
	x[4] = 100
	fmt.Println(x)

	var y [5]float64
	y[0] = 98
	y[1] = 93
	y[2] = 77
	y[3] = 82
	y[4] = 83

	var total float64 = 0
	for i := 0; i < 5; i++ {
		total += y[i]
	}
	fmt.Println(total/5)

	for j := 0; j < len(y); j++ {
		total += y[j]
	}
	fmt.Println(total / float64(len(x)))

	for _, value := range y {
		total += value
	}
	fmt.Println(total / float64(len(x)))

	z := [5]float64{ 98, 92, 77, 82, 83, }
	fmt.Println(z)
}
