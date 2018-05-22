package main

import (
	"fmt"
)

func main() {
	s := "Hello 월드"

	s2 := []rune(s) // UTF-8에서는 한글 및 한자는 1~3 byte 영문은 1byte 글자마다 끊어서 보여주는 rune 배열
	// 다른 프로그래밍 언어들에서는 아스키2 를 사용해 모든 글자가 1byte로 처리됨.
	fmt.Println("len(s2)=", len(s2))

	for i := 0; i < len(s2); i++ {
		fmt.Print(string(s2[i]), ", ")
	}
}
