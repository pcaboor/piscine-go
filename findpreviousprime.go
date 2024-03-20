package piscine

func FindPreviousPrime(nb int) int {
	for i := nb; i >= 2; i-- {
		if IsPrime2(i) {
			return i
		}
	}
	return 0
}

func IsPrime2(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	i := 5

	for i*i <= n {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
		i += 6
	}
	return true
}
