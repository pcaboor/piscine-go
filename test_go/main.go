package main

import (
	"fmt"
	"piscine"
)

func main() {
	a1 := []int{0, 1, 2, 3, 4, 5}
	a2 := []int{0, 2, 1, 3}

	result1 := piscine.IsSorted(f, a1)
	result2 := piscine.IsSorted(f, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}

func f(a, b int) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}
