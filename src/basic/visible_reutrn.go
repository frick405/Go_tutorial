package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
} // Go의 반환 값은 이름이 정해질 수도 있다. 그 경우 함수의 맨 위에서 정의된 변수처럼 다뤄진다.
// 이런 이름들은 반환 값의 의미를 설명ㄷ하는데 사용되며 인자가 없는 return 문은 이름이 주어진 반환 값을 반환하면
// naked return이라고 한다.

func main() {
	fmt.Println(split(17))
}
