package main

import "fmt"

func setPoint(x *int, y *int) {
	*x = 42
	*y = 21
}

func main() {
	var x, y int
	a := &x
	b := &y
	setPoint(a, b)
	fmt.Printf("x = %d, y = %d\n", *a, *b)
}
