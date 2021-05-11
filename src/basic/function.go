package main

import "fmt"

func add(x int, y int) int {
	return x + y
} // 변수 이름 뒤에 type이 온다.

func main() {
	fmt.Println(add(42, 13))
}
