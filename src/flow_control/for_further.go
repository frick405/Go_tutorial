package main

import "fmt"

func main() {
	sum := 1
	for ; sum <1000;{
		sum += sum
	}// 초기화/사후 구문은 필수 아님
	fmt.Println(sum)
}
