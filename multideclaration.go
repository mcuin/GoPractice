package main

import "fmt"

func main() {
	var (
		a = 5
		b = 10
		c = 15
	)
	fmt.Println(a, b, c)
	const (
		d = 20
		e = 25
		f = 30
	)
	fmt.Println(d, e, f)
}
