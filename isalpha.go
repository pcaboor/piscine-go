package piscine

func IsAlpha(s string) bool {
	for _, goodLetter := range s {
		if (goodLetter < 'a' || goodLetter > 'z') && (goodLetter < '0' || goodLetter > '9') && (goodLetter < 'A' || goodLetter > 'Z') {
			return false
		}
	}
	return true
}
