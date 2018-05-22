package main

import (
	"fmt"
)

func add(x, y int) int {
	return x + y
}

func func1(x, y int) (int, int) {
	return y, x
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d + %d = %d\n", i, i+2, add(i, i+2))
	}

	a, b := func1(2, 3)
	fmt.Printf("%d, %d", a, b)
}
