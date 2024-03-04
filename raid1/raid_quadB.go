package piscine

import (
	"github.com/01-edu/z01"
)

func QuadB(x, y int) {
	if x == 1 {
		for j := 1; j <= y; j++ {
			if j == 1 {

				z01.PrintRune('/')
				z01.PrintRune('\n')
			} else if j == y {
				z01.PrintRune(92)
				z01.PrintRune('\n')
			} else {
				z01.PrintRune('*')
				z01.PrintRune('\n')
			}
		}
		return
	}
	for j := 1; j <= y; j++ {
		if j == 1 {
			for i := 1; i <= x; i++ {
				if i == 1 {
					z01.PrintRune('/')
				} else if i == x {
					z01.PrintRune(92)
					z01.PrintRune('\n')
				} else {
					z01.PrintRune('*')
				}
			}
		} else if j == y {
			for i := 1; i <= x; i++ {
				if i == 1 {
					z01.PrintRune(92)
				} else if i == x {
					z01.PrintRune('/')
					z01.PrintRune('\n')
				} else {
					z01.PrintRune('*')
				}
			}
		} else {
			for i := 1; i <= x; i++ {
				if i == 1 {
					z01.PrintRune('*')
				} else if i == x {
					z01.PrintRune('*')
					z01.PrintRune('\n')
				} else {
					z01.PrintRune(' ')
				}
			}
		}
	}
}
