package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	var v float64 = math.Pow(x, n)
	if v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	} // else cluase shuld be after if clause

	return lim
}

func main() {
	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))
}

