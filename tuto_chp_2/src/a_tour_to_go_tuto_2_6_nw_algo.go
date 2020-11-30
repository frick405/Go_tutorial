package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	var z float64 = 2

	for i := 0; i < 10; i++ {
			if z == math.Sqrt(x) {
					return z
			} else {
					z -= (z*z - x) / (2*z)
			}
		}
	return z
}

func main() {
		fmt.Println(Sqrt(4))
		fmt.Println(math.Sqrt(4))
}


