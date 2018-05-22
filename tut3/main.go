package main

import "fmt"

//구구단 생성
func main() {

	for dan := 1; dan <= 9; dan++ {

		// 5단을 건너뛰고 구구단 생성시
		if dan == 5 {
			continue
		}
		fmt.Printf("%d단\n", dan)

		for j := 1; j <= 9; j++ {
			// 각 단의 5를 건너뛰고 구구단 생성시
			if j == 5 {
				continue
			}
			fmt.Printf("%d * %d = %d \n", dan, j, dan*j)
		}
		fmt.Println()
	}
}
