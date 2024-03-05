package piscine

func IsPrime(nb int) bool {
	if nb == 1 {
		return false
	}
	if nb == 0 {
		return false
	}
	if nb%2 != 0 {
		return true
	}
	if nb < 0 || nb%2 != 0 {
		return false
	}
	return false
}

// nombre entier nb divisble par 1 et par lui même
