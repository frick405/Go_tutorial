// In Golang, Function can get 0 or more paramete

package main

import "fmt"
/*
// Should declare data type!
func add(x int, y int) int {
	return x + y
}*/

// If parameters have same data type, Then we can omit data type declaration except for last paramter
func add(x, y int) int {
	return x + y
}


func main() {
	fmt.Println(add(12, 24))
}
