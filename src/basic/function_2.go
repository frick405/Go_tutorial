package main

import "fmt"

func add(x, y int) int {
	return x + y
} // 주어진 매개변수가 같은 타입이면 생략 가능

func main() {
	fmt.Println(add(42, 13))
}
