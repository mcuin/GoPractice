package main

import "fmt"
import "math"

type Circle struct{
	x float64;
	y float64;
	z float64;
}



func area(c *Circle) float64 {
	return math.Pi * c.z * c.z
}

func main() {
	c := Circle{0, 0, 5}
	fmt.Println(area(&c))
}
