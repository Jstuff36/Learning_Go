// This is my first go problem

package main

import (
	"fmt"
)

func Sqrt(num float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - num) / (2 *  z)	
		fmt.Println(z)
	}
	return z
}

func SqrtSmartStop(num float64) float64 {
	curr := 1.0
	old := 0.0
	for old - curr != 0 {
		old = curr
		curr -= (curr * curr - num) / (2 * curr)
		fmt.Println(curr)
	}
	return curr
}

func main() {
	fmt.Println(SqrtSmartStop(1424))
}