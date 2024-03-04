package piscine

import (
	"github.com/01-edu/z01"
)

func QuadD(x, y int) {
	if x == 1 {
		for j := 1; j <= y; j++ {
			if j == 1 {
				z01.PrintRune('A')
				z01.PrintRune('\n')
			} else if j == y {
				z01.PrintRune('C')
				z01.PrintRune('\n')
			} else {
				z01.PrintRune('B')
				z01.PrintRune('\n')
			}
		}
		return
	}
	for j := 1; j <= y; j++ {
		if j == 1 {
			for i := 1; i <= x; i++ {
				if i == 1 {
					z01.PrintRune('A')
				} else if i == x {
					z01.PrintRune('A')
					z01.PrintRune('\n')
				} else {
					z01.PrintRune('B')
				}
			}
		} else if j == y {
			for i := 1; i <= x; i++ {
				if i == 1 {
					z01.PrintRune('A')
				} else if i == x {
					z01.PrintRune('C')
					z01.PrintRune('\n')
				} else {
					z01.PrintRune('B')
				}
			}
		} else {
			for i := 1; i <= x; i++ {
				if i == 1 {
					z01.PrintRune('B')
				} else if i == x {
					z01.PrintRune('B')
					z01.PrintRune('\n')
				} else {
					z01.PrintRune(' ')
				}
			}
		}
	}
}
