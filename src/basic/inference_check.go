package main

import "fmt"

func main() {
	v := 42 // := or var표현을 이요해 명시적인 type을 정의하지 않고 변숙를 선언할 때 그 변수 type은 오른편에 있는 값으로부터 유추된다.
	fmt.Printf("v is of type %T\n", v)
}
