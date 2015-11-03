// Basics
package main

import (
	"fmt"
	"math"
)

var (
	name     string = "Prince Obyron"
	location string = "Dorne"
	age      int    = 32

	name2, location2, age2 = "Prince", "Dorne", 32
)

const (
	StatusOK                 = 200
	StatusCreated            = 201
	StatusAccepted           = 202
	StatusNonAuthorativeInfo = 203
	StatusNoContent          = 204
	StatusResetContent       = 205
	StatusPartialContent     = 206
	Pi                       = 3.14
	Truth                    = false
	Big                      = 1 << 100
	Small                    = Big >> 99
)

func main() {
	name3, location3 := "Prince Obyron", "Dorne"
	age3 := 32
	fmt.Printf("%s (%d) of %s\n", name3, age3, location3)

	const greeting = "ハローワールド"
	fmt.Println(greeting)
	fmt.Println(Pi)
	fmt.Println(Truth)

	fmt.Println(math.Pi)
}
