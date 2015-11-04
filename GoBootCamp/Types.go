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
)

func timeMap(y interface{}) {
	z, ok := y.(map[string]interface{})
	if ok {
		z["updated at"] = time.Now()
	}
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
}
