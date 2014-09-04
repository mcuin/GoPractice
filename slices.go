package main

import "fmt"

func main() {
	//var x []float64
	//y := make([]float64, 5)
	//z := make([]float64, 5, 10)
	//arr := [5]float64{1,2,3,4,5,}
	//a := arr[0:5]
	slice1 := []int{1,2,3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)

	slice3 := []int{1, 2, 3}
	slice4 := make([]int, 2)
	copy(slice4, slice3)
	fmt.Println(slice3, slice4)
}
