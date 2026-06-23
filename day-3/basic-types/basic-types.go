package main

import (
	"fmt"
)

var (
	boole  bool       = false
	Maxint uint64     = 1<<64 - 1
	z      complex128 = (-5 + 13i)
)

func main() {
	fmt.Printf("Type -: %T Value -: %v ", boole, boole)
	fmt.Printf("Type -: %T Value -: %v ", Maxint, Maxint)
	fmt.Printf("Type -: %T Value -: %v ", z, z)
}
