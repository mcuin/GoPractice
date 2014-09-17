package main

import ("fmt" 
	"time" 
	"math/rand")

func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func printer(c chan string) {
	for {
		msg := <- c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

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

	var c chan string = make(chan string)
	
	go pinger(c)
	go ponger(c)
	go printer(c)

	var input4 string
	fmt.Scanln(&input4)

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}() 
	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()
	go func() {
		for {
			select {
				case msg1 := <- c1: fmt.Println("Message 1", msg1)
				case msg2 := <- c2: fmt.Println("Message 2", msg2)
				case <- time.After(time.Second): fmt.Println("timeout")
				default: fmt.Println("Nothing ready.")
			}
		}
	}()

	var input5 string
	fmt.Scanln(&input5)
}
