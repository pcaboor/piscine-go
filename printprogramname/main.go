package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) > 0 {
		for _, r := range os.Args[0] {
			z01.PrintRune(r)
		}
	}
}
