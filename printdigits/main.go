package main

import (
	"github.com/01-edu/z01"
)

func main() {

	var numOne rune
	for numOne = 48; numOne <= 57; numOne++ {
		z01.PrintRune(numOne)
	}
	z01.PrintRune('\n')
}
