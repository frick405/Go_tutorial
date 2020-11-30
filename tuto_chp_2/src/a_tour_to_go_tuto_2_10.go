// Deferred function called is stacked in stack
// If one function ends, deferred functions work LIFO

package main

import (
	"fmt"
)

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
		}
}

