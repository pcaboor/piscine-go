package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb2() {
	for w := '0'; w <= '9'; w++ {
		for x := '0'; x <= '9'; x++ {
			for y := '0'; y <= '9'; y++ {
				for z := '0'; z <= '9'; z++ {
					if w < y || (w == y && x < z) {
						z01.PrintRune(w)
						z01.PrintRune(x)
						z01.PrintRune(' ')
						z01.PrintRune(y)
						z01.PrintRune(z)
						if w != '9' || x != '8' || y != '9' || z != '9' {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
