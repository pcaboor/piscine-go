package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	for j := 0; j < len(a)-1; j++ {
		if f(a[j], a[j+1]) > 0 {
			return false
		}
	}
	return true
}

func f(a, b int) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}
