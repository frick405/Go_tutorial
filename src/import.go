package main

import (
	"fmt"
	"math"
) // Factored import 문이라고 불림, 따로 import 할 수도 있지만 Factored import가 권장됨

func main() {
	fmt.Printf("Now you have %g problems. \n", math.Sqrt(7))
}
