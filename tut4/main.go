package main

import "fmt"

// 별 생성
func main() {

	// 증가하는 부분
	for i := 0; i < 4; i++ {
		// 공백 부분
		for j := 0; j < 3-i; j++ {
			fmt.Print(" ")
		}
		// 별 부분
		for j := 0; j < 2*i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	// 감소하는 부분
	for i := 0; i < 3; i++ {
		//공백 부분
		for j := 0; j < i+1; j++ {
			fmt.Print(" ")
		}
		//별 부분
		for j := 0; j < 5-2*i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
