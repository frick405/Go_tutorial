package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {

	var v float64 = math.Pow(x, n)
	// v's scope is in if clause. Out of if claouse, v is nothing 
	if v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(pow(3, 2, 10),pow(3, 3, 20))
}



