package piscine

func IsLower(s string) bool {
	for _, goodLetter := range s {
		if goodLetter < 'a' || goodLetter > 'z' {
			return false
		}
	}
	return true
}
