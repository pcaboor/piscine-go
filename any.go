package piscine

func Any(f func(string) bool, a []string) bool {
	for _, n := range a {
		if f(n) {
			// vérifie si la chaine est composée de chiffres
			return true
		}
	}
	return false
}

func IsNumericNumb(s string) bool {
	for _, goodLetter := range s {
		if goodLetter < '0' || goodLetter > '9' {
			return false
		}
	}
	return true
}
