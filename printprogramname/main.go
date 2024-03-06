package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	a := os.Args[0]
	for _, r := range a {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
