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
