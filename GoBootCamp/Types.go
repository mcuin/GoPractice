// Types
package main

import (
	"fmt"
	"math/cmplx"
	"time"
)

var (
	goIsFun        = true
	maxInt  uint64 = 1<<64 - 1
	complex        = cmplx.Sqrt(-5 + 12i)

	p = Point{1, 2}
	q = &Point{1, 2}
	r = Point{X: 1}
	t = Point{}
)

func timeMap(y interface{}) {
	z, ok := y.(map[string]interface{})
	if ok {
		z["updated at"] = time.Now()
	}
}

type Stringer interface {
	String() string
}

type fakeString struct {
	content string
}

func (s *fakeString) String() string {
	return s.content
}

func printString(value interface{}) {
	switch str := value.(type) {
	case string:
		fmt.Println(str)
	case Stringer:
		fmt.Println(str.String())
	}
}

type BootCamp struct {
	Lat  float64
	Lon  float64
	Date time.Time
}

type Point struct {
	X, Y int
}

func main2() {
	const f = "%T(%v)\n"
	fmt.Printf(f, goIsFun, goIsFun)
	fmt.Printf(f, maxInt, maxInt)
	fmt.Printf(f, complex, complex)

	integer := 42
	float := float64(integer)
	unsigned := uint(float)

	fmt.Println(unsigned)

	foo := map[string]interface{}{
		"Matt": 42,
	}

	timeMap(foo)
	fmt.Println(foo)

	s := &fakeString{"Ceci n'est pas un string"}
	printString(s)
	printString("Hello gophers")

	fmt.Println(BootCamp{
		Lat:  34.012836,
		Lon:  -118.495338,
		Date: time.Now(),
	})

	fmt.Println(p, q, r, t)

	event := BootCamp{
		Lat: 34.012836,
		Lon: -118.495338,
	}

	event.Date = time.Now()
	fmt.Printf("Event on %s, location (%f, %f)\n", event.Date, event.Lat, event.Lon)
}
