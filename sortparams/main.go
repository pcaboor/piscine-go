package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:] // Exclure le nom du programme
	for i := 0; i <= 127; i++ {
		for _, ch := range args {
			if len(ch) > 0 && int(ch[0]) == i {
				os.Stderr.WriteString(ch)
				z01.PrintRune('\n')
			}
		}
	}
}
