package main

import (
	"fmt"
)

func main() {
		var i, j int = 1, 2
		k := 3 // := can be used like var declaration in function. Out of function, then we can't use that.
		c, python, java := true, false, "no!"

		fmt.Println(i, j, k, c, python, java)
}
