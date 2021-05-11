package main

import (
	"fmt"
	"math"
)
// if문 또한 조건문 전에 수행될 잛은 구문으로 시작될 수 있다.
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}

	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20)
	)
}
