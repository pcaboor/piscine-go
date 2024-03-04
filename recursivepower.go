package piscine

func RecursivePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}
	if nb == 0 || nb == 1 || power == 1 {
		return nb
	}
	return nb * RecursivePower(nb, power-1)
}
