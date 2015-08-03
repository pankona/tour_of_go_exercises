package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	preResult := z
	for {
		z = z - ((z*z - x) / (2 * x))
		if math.Abs(preResult-z) < 0.00000001 {
			return z
		} else {
			preResult = z
		}
	}
}

func main() {
	fmt.Println(Sqrt(2))
}
