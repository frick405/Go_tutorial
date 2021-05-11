package main

import "fmt"

const Pi = 3.14


func main() {
	const World = "세계"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true // 상수는 변수처럼 선언되며 string, boolean, 숫자 값이 될 수 있습니다.
	fmt.Println("Go rules?", Truth)
}
