// Basics
package main

import (
	"fmt"
)

var (
	name     string = "Prince Obyron"
	location string = "Dorne"
	age      int    = 32

	name2, location2, age2 = "Prince", "Dorne", 32
)

func main() {
	name3, location3 := "Prince Obyron", "Dorne"
	age3 := 32
	fmt.Printf("%s (%d) of %s", name3, age3, location3)
}
