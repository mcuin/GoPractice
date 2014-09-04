package main

import "fmt"

func main() {
	var a string = "Hello World"
	fmt.Println(a)

	var b string
	b = "Hello World"
	fmt.Println(b)

	var c string
	c = "first"
	fmt.Println(c)
	c = "second"
	fmt.Println(c)

	var d string
	d = "first"
	fmt.Println(d)
	d = d + "second"
	fmt.Println(d)

	var e string = "hello"
	var f string = "world"
	fmt.Println(e == f)

	var g string = "hello"
	var h string = "hello"
	fmt.Println(g == h)

	i := "Hello World"
	fmt.Println(i)
	var j = "Hello World"
	fmt.Println(j)
	k := 5
	fmt.Println(k) 
}
