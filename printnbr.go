package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	var str string

	if n == 0 {
		z01.PrintRune('0')
	}
	if n < 0 {
		str = "-"
		n = -n
	}
	for n != 0 {
		digit := n % 10
		str = string(rune(digit+'0')) + str
		n /= 10
	}
	for _, s := range str {
		z01.PrintRune(s)
	}
}
