package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	if n == 0 {
		z01.PrintRune('0')
	}
	if n > 0 {
		digit := n % 10
		z01.PrintRune(rune(digit + '0'))
		n /= 10
	}
}
