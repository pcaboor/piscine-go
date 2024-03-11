package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	if n <= 0 {
		z01.PrintRune('0')
		return
	}
	if n != 0 {
		var t []int
		var swap int
		stock := 0
		countN := 0

		for n > 0 {
			stock = n % 10
			n /= 10
			t = append(t, stock)
		}

		for count := range t {
			countN = count + 1
		}
		for i := 0; i < countN-1; i++ {
			for j := 0; j < countN-i-1; j++ {
				if t[j] > t[j+1] {
					swap = t[j]
					t[j] = t[j+1]
					t[j+1] = swap
				}
			}
		}
		for k := 0; k < countN; k++ {
			z01.PrintRune(rune(t[k] + 48))
		}
	}
}
