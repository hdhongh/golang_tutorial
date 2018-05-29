package main

import (
	"fmt"
)

func removeBack(a []int) ([]int, int) {
	return a[:len(a)-1], a[len(a)-1]
}

func removeFront(a []int) ([]int, int) {
	return a[1:], a[0]
}

func main() {

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < 5; i++ {
		//var back int
		//a, back = removeBack(a)
		//fmt.Printf("%d, ", back)
		var front int
		a, front = removeFront(a)
		fmt.Printf("%d, ", front)
	}
	fmt.Println()
	fmt.Println(a)

}
