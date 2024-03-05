package piscine

func IsNumeric(s string) bool {
	for _, goodLetter := range s {
		if goodLetter < '0' || goodLetter > '9' {
			return false
		}
	}
	return true
}
