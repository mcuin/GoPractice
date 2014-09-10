package main

import "fmt"
//import "os"

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
} 

func mult() (int, int) {
	return 5, 6
}

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func evens() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	
	return x * factorial(x - 1)
}

func first() {
	fmt.Println("First")
}

func second(){
	fmt.Println("Second")
}

func main() {
	xs := []float64{98, 93, 77, 82, 83}
	fmt.Println(average(xs))

	someOther := []float64{98, 93, 77, 82, 83,}
	fmt.Println(average(someOther))

	x, y := mult()
	fmt.Println(x, y)

	fmt.Println(add(1,2,3))

	addOther := []int{1,2,3}
	fmt.Println(add(addOther...))

	z := 0
	increment := func() int {
		z++
		return z
	}
	fmt.Println(increment())
	fmt.Println(increment())

	nextEven := evens()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())	

	fmt.Println(factorial(4))
	
	defer second()
	first()

	/*f, _ := os.Open(filename)
	defer f.Close()*/ 
}
