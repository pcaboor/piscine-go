package main

import (
	"github.com/01-edu/z01"
)

func main() {
	var rRune rune
	for rRune = 122; rRune >= 97; rRune-- {
		z01.PrintRune(rRune)
	}
	z01.PrintRune('\n')
}
