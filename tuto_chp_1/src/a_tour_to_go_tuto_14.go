package main

import (
	"fmt"
)

var(
	big = 1 << 30
	small = big >> 29
)

func needInt(x int) {
		return (x*10 - 1)
}

func needFloat(x float64) {
		return (x*10)
}

func main() {
		// Create huge number from 1 bit by shifting 100 places
		fmt.Println(big)
		fmt.Println(small)
		//fmt.Println(needInt(big))
		//fmt.Println(needInt(big))
}

