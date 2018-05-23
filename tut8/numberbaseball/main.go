package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 숫자 야구 게임
// 두 사람
// 각 사람 0-9 까지 겹치지 않는 숫자
// 숫자는 같고 자릿수가 다르면 1Ball
// 숫자도 같고 자릿수도 같으면 1Strike
// 게임을 진행하면서 상대 숫자 추측하는 게임
// 3S 가 될 때까지 진행

func MakeNumbers() [3]int {
	//0~9 사이의 겹치지 않는 무작위 숫자 3개를 반환한다.

	// ntn returns, as an int, a non-negative pseudo-random number in [0,n) from the default Source
	// [0,n) 은 n을 포함하지 않는 n-1까지라는 뜻.
	var rst [3]int
	for i := 0; i < 3; i++ {
		for {
			n := rand.Intn(10)
			dup := false
			for j := 0; j < i; j++ {
				if rst[j] == n {
					//자신의 앞의 수와 비교 겹친다 다시 뽑는다.
					dup = true
					break
				}
			}
			// 겹치지 않을 경우에 break
			if !dup {
				rst[i] = n
				break
			}
		}
	}
	fmt.Println(rst)
	return rst
}

func InputNumbers() [3]int {
	// 키보드로 부터 0~9 사이의 겹치지 않는 무작위 숫자 3개를 입력받아 반환한다.
	var rst [3]int

	for {
		var no int
		_, err := fmt.Scanf("%d\n", &no)

		if err != nil {
			fmt.Println("잘못 입력하셨습니다.")
			continue
		}

		idx := 2
		for no > 0 {
			n := no % 10
			no = no / 10
			// 겹치지 않는 숫자 입력 검증 필요
			rst[idx] = n
			idx--
		}
		break
	}

	fmt.Println(rst)

	return rst
}

func CompareNumbers(numbers, inputNumbers [3]int) bool {
	// 두개의 숫자 3개를 비교해서 결과를 반환한다.
	return true
}

func PrintResult(result bool) {
	fmt.Println(result)
}

func IsGameEnd(result bool) bool {
	// 비교 결과가 3 스트라이크인지 확인
	return true
}

func main() {
	// 랜덤숫자를 위해 Seed값 설정
	rand.Seed(time.Now().UnixNano())
	// 무작위 숫자 3개를 만든다.
	numbers := MakeNumbers()

	cnt := 0
	for {
		cnt++
		// 사용자의 입력을 받는다.
		inputNumbers := InputNumbers()

		// 결과를 비교한다.
		result := CompareNumbers(numbers, inputNumbers)

		// Result Print한다
		PrintResult(result)

		// 게임이 3S인가?
		if IsGameEnd(result) {
			//게임끝
			break
		}
	}

	// 게임끝. 몇 번 만에 맞췄는지 출력
	fmt.Printf("%d 번 만에 맞췄다.\n", cnt)
}
