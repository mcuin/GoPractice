package main

import "fmt"

var x string = "Hello World"

func main() {
	fmt.Println(x)
	foo()
}

func foo(){
	fmt.Println(x)
}
