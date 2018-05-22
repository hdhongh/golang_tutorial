package main

import (
	"fmt"
)

func main() {
	var a int // 선어문
	a = 4     // 대입문

	b := 4 // 선언대입문

	var c = 6 // TYPE 선언 없이 사용하기

	fmt.Printf("%d\n", a) // %d decimal
	fmt.Printf("%v\n", b) // %v value (타입은 알아서 정의됨)
	fmt.Printf("%v\n", c)

	operator() // 이항연산자 테스트
}

func operator() {
	a := 4
	b := 2

	fmt.Printf("a&b = %v\n", a&b)
	fmt.Printf("a|b = %v\n", a|b)
	fmt.Printf("a/b = %v\n", a/b)
	fmt.Printf("ab = %v\n", a%b)
}
