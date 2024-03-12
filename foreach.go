package piscine

import "github.com/01-edu/z01"

func ForEach(f func(int), a []int) {
	for _, n := range a {
		f(n)
	}
	z01.PrintRune('\n')
}
