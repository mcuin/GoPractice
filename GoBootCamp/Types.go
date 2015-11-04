// Types
package main

import (
	"fmt"
	"math/cmplx"
)

var (
	goIsFun        = true
	maxInt  uint64 = 1<<64 - 1
	complex        = cmplx.Sqrt(-5 + 12i)
)

func main2() {
	const f = "%T(%v)\n"
	fmt.Printf(f, goIsFun, goIsFun)
	fmt.Printf(f, maxInt, maxInt)
	fmt.Printf(f, complex, complex)
}
