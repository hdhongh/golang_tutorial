package main

import (
	"fmt"
)

func main() {
	//s := sum(10, 0)
	s := sumFor(10)
	fmt.Println(s)
}

// 1~10 까지 합계를 출력
func sum(x, s int) int {
	// 탈출문이 꼭 필요
	if x == 0 {
		return s
	}
	s += x
	return sum(x-1, s)
}

// 모든 재귀호출 구문은 반복문으로 바꿀 수 있음.
func sumFor(x int) int {
	var s int
	for i := 0; i <= x; i++ {
		s += i
	}
	return s
}
