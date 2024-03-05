package piscine

func IsPrintable(s string) bool {
	for _, goodLetter := range s {
		if goodLetter < ' ' || goodLetter > '~' {
			return false
		}
	}
	return true
}

/* func IsPrintable(s string) bool {
	var aRune rune
	for _, goodLetter := range s {
		if goodLetter < 32 || goodLetter > 127 {
			return false
		}
	}
	return true
}
*/
