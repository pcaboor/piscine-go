package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	lengthOfArg := len(os.Args) - 1
	EvenMsg := "I have an odd number of arguments"
	OddMsg := "I have an even number of arguments"

	if isEven(lengthOfArg) {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if nbr == 1 {
		return true
	} else {
		return false
	}
}
