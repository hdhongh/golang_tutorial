package main

import (
	"bufio" // 입력으로부터 한줄 입력 받기 위해
	"fmt"
	"os" // 표준입력
	"strconv"
	"strings" // 문자열 찌꺼기 제거
	// 문자열 -> 숫자열
)

// 계산기 작성
func main() {
	fmt.Println("숫자를 입력하세요")
	reader := bufio.NewReader(os.Stdin) // os 의 Stdin 을 bufio의 새로운 reader 객체에 넣기
	line, _ := reader.ReadString('\n')  // \n 이 나올 때까지 읽기
	line = strings.TrimSpace(line)      // 공백제거

	n1, _ := strconv.Atoi(line) // 문자열 -> Integer

	line, _ = reader.ReadString('\n') // \n 이 나올 때까지 읽기
	line = strings.TrimSpace(line)    // 공백제거

	n2, _ := strconv.Atoi(line) // 문자열 -> Integer

	fmt.Printf("입력하신 숫자는 %d, %d 입니다.", n1, n2)

	fmt.Println("연산자를 입력하세요")

	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)

	if line == "+" {
		fmt.Printf("%d + %d = %d\n", n1, n2, n1+n2)
	} else if line == "-" {
		fmt.Printf("%d - %d = %d\n", n1, n2, n1-n2)
	} else if line == "*" {
		fmt.Printf("%d * %d = %d\n", n1, n2, n1*n2)
	} else if line == "/" {
		fmt.Printf("%d / %d = %d\n", n1, n2, n1/n2)
	} else {
		fmt.Println("잘못 입력하셨습니다.")
	}
}
