package piscine

func RecursiveFactorial(nb int) int {

	if nb < 0 || nb > 20 {
		return 0
	}
	if nb == 0 || nb == 1 {
		return 1
	}
	for i := 1; i <= nb; i++ {

	}
	return nb * RecursiveFactorial(nb-1)
}
