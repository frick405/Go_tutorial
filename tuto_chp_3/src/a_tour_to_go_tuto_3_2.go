package main

import (
	"fmt"
)
// Struct is set of field
// Same like python class
type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}

