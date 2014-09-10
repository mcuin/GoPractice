package main

import "fmt"

func zero(xPtr *int) {
	*xPtr = 0
}

func one(yPtr *int) {
	*yPtr = 1
}

func main() {
	x := 5
	zero(&x)
	fmt.Println(x)

	yPtr := new(int)
	one(yPtr)
	fmt.Println(*yPtr)
}
