package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000{ // 세미콜론을 삭제하고 for를 사용할 수 있다. 
		sum += sum
	}
	fmt.Println(sum)
}
