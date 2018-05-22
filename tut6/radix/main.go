package main

import (
	"fmt"
)

// radix sort O(N)성능
// 들어가는 원소의 범위가 제한되어야 한다.
// 들어가는 원소의 범위가 작아야 한다.
// 원소의 범위 만큼 배열을 하나 생성해서 각 인덱스마다 해당하능 원소의 카운트를 적립.
func main() {
	arr := [10]int{0, 5, 4, 2, 8, 4, 3, 4, 5, 6}
	temp := [10]int{}

	for i := 0; i < len(arr); i++ {
		idx := arr[i] // arr의 값을 idx 에 저장
		temp[idx]++   // 해당 idx 의 카운트 적립
	}

	idx := 0
	for i := 0; i < len(temp); i++ {
		for j := 0; j < temp[i]; j++ { // 적립된 카운트 만큼 반복하면서 i 값 arr에 대입
			arr[idx] = i
			idx++
		}
	}
	fmt.Println("radix result = ", arr)
}
