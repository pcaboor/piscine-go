package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	var str string

	if n < 0 {
		n = -n
		z01.PrintRune('-')
	} else if n == 0 {
		z01.PrintRune('0')
	}

	for n > 0 {
		numb := n % 10
		str = string(rune(numb+'0')) + str
		n /= 10
	}

	for _, s := range str {
		z01.PrintRune(s)
	}
}
