package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	// After condition, you should use {}!
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(4), sqrt(-4))
}

