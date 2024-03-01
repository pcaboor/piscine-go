package piscine

import (
	"github.com/01-edu/z01"
)

func PrintStr(s string) {
	for _, myHelloWord := range s { //_, stocker de manière temporaire et range "dans mon tableau"
		z01.PrintRune(myHelloWord)
	}
}

/* ou sans rune

func StrRev(s string) string {
	aString := []byte(s)
	return string(aString)
}
*/
