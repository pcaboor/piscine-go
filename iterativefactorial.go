package piscine

func IterativeFactorial(nb int) int {
	result := 1

	if nb < 0 || nb > 20 {
		return 0
	}
	if nb == 0 || nb == 1 {
		return 1
	}

	for i := 1; i <= nb; i++ {

		result = result * i

	}
	return result
}
