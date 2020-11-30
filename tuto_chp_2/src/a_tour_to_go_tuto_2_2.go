package main

import (
	"fmt"
)

var sum int = 1

// And semicolon can be omitted
// If you want to infinity loop, then clean your condition clause
func main() {
	for ; sum < 10; {
		sum += sum
	}
	fmt.Println(sum)
}
