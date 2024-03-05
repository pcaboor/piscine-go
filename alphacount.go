package piscine

func AlphaCount(s string) int {
	count := 0
	for _, countLetter := range s {
		if (countLetter >= 'A' && countLetter <= 'Z') || (countLetter >= 'a' && countLetter <= 'z') {
			count++
		}
	}
	return count
}
