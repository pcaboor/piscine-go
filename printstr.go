package piscine

import (
	"github.com/01-edu/z01"
)

func PrintStr(myHelloWord string) {
	for _, myHelloWord := range myHelloWord { //_, stocker de manière temporaire et range "dans mon tableau"
		z01.PrintRune(myHelloWord)
	}
}

/* ou sans rune

func StrRev(s string) string {
	aString := []byte(s)
	return string(aString)
}
*/
