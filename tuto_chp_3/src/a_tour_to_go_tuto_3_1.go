package main

import (
	"fmt"
)

func main() {
	var i, j int = 42, 2701

	p := &i // pointer to i
	fmt.Println(p) // address which is same for i
	fmt.Println(*p) // value in address which is same for i
	
	*p = 21 // new value is assigned to i pointer
	fmt.Println(p) 
	p = &j //new pointer to j
	fmt.Println(p) // j's pointer is assgined to p
	fmt.Println(*p) // value in j pointer
	fmt.Println(i)
	fmt.Println(j)
}
