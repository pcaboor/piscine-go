package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	if n > 0 {
		var t []int
		tCount := 0
		v := 0
		var fV int
		for n != 0 {
			v = n % 10
			n /= 10
			t = append(t, v)
		}
		for count := range t {
			tCount = count + 1
		}
		for i := 0; i < tCount-1; i++ {
			for j := 0; j < tCount-i-1; j++ {
				if t[j] > t[j+1] {
					fV = t[j]
					t[j] = t[j+1]
					t[j+1] = fV
				}
			}
		}
		for k := 0; k < tCount; k++ {
			z01.PrintRune(rune(t[k] + 48))
		}
	}

}
