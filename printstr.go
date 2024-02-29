package piscine

import (
	"github.com/01-edu/z01"
)

func PrintStr(myHelloWord string) {
	for _, myHelloWord := range myHelloWord {
		z01.PrintRune(myHelloWord)
	}
}
