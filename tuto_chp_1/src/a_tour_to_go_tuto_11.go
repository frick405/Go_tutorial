// T(value) can convert original value data type to new data type
package main

import (
	"fmt"
	"math"
)

func main () {
		var x, y int = 3, 5
		var f float64 = math.Sqrt(float64(x*x + y*y)) // from float?? to float6
		var f_2 int = int(math.Sqrt(float64(x*x + y*y)))
		var z uint = uint(f) //from float64 to uint // uint type is unsigend(only positive number), uint range is double in aspect of positive number
		fmt.Println(x, y, z, f, f_2)
}
