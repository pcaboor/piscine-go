package piscine

func IterativeFactorial(nb int) int {
	result := 0

	for i := 0; i <= nb+1; i++ {
		if nb < 0 || nb > 12 {
			return 0
		}
		if nb == 0 || nb == 1 {
			return 1
		} else {
			result = result + nb
		}
	}
	return result
}
