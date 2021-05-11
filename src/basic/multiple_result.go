package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
} // 한 함수는 몇개의 결과든 반환할 수 있다.

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
