package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:] // Exclure le nom du programme
	for i := 0; i <= 127; i++ {
		for _, arg := range args {
			if len(arg) > 0 && int(arg[0]) == i {
				os.Stderr.WriteString(arg)
				z01.PrintRune('\n')
			}
		}
	}
}
