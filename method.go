package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hello my name is", p.Name)
}

type Android struct {
	Person
	model string
}

func main() {
	a := new(Android)
	a.Talk()
}


