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

func add(x, y int) int {
	return x + y
}

func local(city string) (string, string) {
	var region string
	var continent string

	switch city {
	case "Los Angeles", "LA", "Santa Monica":
		region, continent = "California", "North America"
	case "New York City", "New York", "NYC":
		region, continent = "New York", "North America"
	default:
		region, continent = "Unknown", "Unknown"
	}

	return region, continent
}

func main() {
	name3, location3 := "Prince Obyron", "Dorne"
	age3 := 32
	fmt.Printf("%s (%d) of %s\n", name3, age3, location3)

	const greeting = "ハローワールド"
	fmt.Println(greeting)
	fmt.Println(Pi)
	fmt.Println(Truth)

	fmt.Println(math.Pi)

	fmt.Println(add(42, 13))

	region, continent := local("Santa Monica")
	fmt.Printf("Matt lives in %s, %s", region, continent)
}
