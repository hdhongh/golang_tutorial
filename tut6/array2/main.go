package main

import (
	"fmt"
)

func main() {
	func2()
}

// 배열 역순으로 만들기 (임시 변수 만들어서 두번 복사 필요)
func func1() {
	arr := [5]int{1, 2, 3, 4, 5}
	temp := [5]int{}

	for i := 0; i < 5; i++ {
		temp[i] = arr[len(arr)-1-i]
	}

	for i := 0; i < 5; i++ {
		arr[i] = temp[i]
	}
	fmt.Println("temp=", temp)
	fmt.Println("arr=", arr)
}

// 배열 역순으로 만들기 (복사 하지 않고 자리만 바꾼다)
func func2() {
	arr := [5]int{1, 2, 3, 4, 5}

	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
	fmt.Println("arr = ", arr)
}
