package main

import (
	"fmt"
)
// defer clause delay of excution till clause containing itself ends
func main() {
	defer fmt.Println("world")
	fmt.Println("Hello")
}

