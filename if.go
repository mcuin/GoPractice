package main

import "fmt"

func main() {
	for j := 1; j <= 10; j++ {
		if j % 2 == 0 {
			fmt.Println(j, "even")
		} else {
			fmt.Println(j, "odd")
		}
	}
}
