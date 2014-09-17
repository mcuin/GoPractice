package main

import ("fmt" 
	"time" 
	"math/rand")

func f2(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}

func main() {
	go f(0)
	var input string
	fmt.Scanln(&input)

	for i := 0; i < 10; i++ {
		go f(i)
	}
	var input2 string
	fmt.Scanln(&input2)

	for i := 0; i < 10; i++ {
		go f2(i)
	}
	var input3 string
	fmt.Scanln(&input3)
}
