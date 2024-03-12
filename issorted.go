package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	for j := 0; j < len(a)-1; j++ {
		if a[j] > a[j+1] {
			return false
		}
	}
	return true
}
