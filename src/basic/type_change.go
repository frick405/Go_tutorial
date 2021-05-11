package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z) // T(v)는 v라는 값을 T type으로 변환시켜 준다
	// go는 다른 type의 요소들 간의 할당에는 명시적인 변환을 필요로 한다.
}

