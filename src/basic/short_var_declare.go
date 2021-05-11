package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!" // := 는 암시적 type으로 var 선언처럼 사용될 수 있음
	// 함수 밖에서는 모든 선언이 키워드로 시작하므로 := 구문은 사용할 수 없습니다.

	fmt.Println(i, j, k, c, python, java)
}
