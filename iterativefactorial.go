package piscine

func IterativeFactorial(nb int) int {
	result := 0
	for i := 0; i <= nb+1; i++ {
		result = result + nb
	}
	return result
}
