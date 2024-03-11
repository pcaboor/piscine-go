package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	ponc := 1

	if n == 0 {
		z01.PrintRune('0')
		return
	}
	if n < 0 {
		ponc = -1
		z01.PrintRune('-')
	}
	if n != 0 {
		a := (n / 10) * ponc
		if a != 0 {
			PrintNbr(a)
		}
		numb := (n % 10 * ponc) + '0'
		z01.PrintRune(rune(numb))
	}
}
