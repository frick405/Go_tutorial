package main

import (
	"fmt"
)
// semicolon works to device each cluase to do their job
// Before ; => Executed before first iteration
// condition clause is next to it. It is used for decision for every iteration
// After }, last clause is executed
func main() {
	var sum int = 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
