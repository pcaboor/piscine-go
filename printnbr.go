package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	var tab [20]int
	index := 0
	for n > 0 {
		tab[index] = n % 10
		n /= 10
		index++
	}
	for r := index - 1; r >= 0; r-- {
		z01.PrintRune(rune('0' + tab[r]))
	}

}
