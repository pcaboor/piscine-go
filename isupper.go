package piscine

func IsUpper(s string) bool {
	for _, goodLetter := range s {
		if goodLetter < 'A' || goodLetter > 'Z' {
			return false
		}
	}
	return true
}
