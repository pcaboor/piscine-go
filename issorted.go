package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	if f(a[0], a[1]) < 0 {
		for j := 0; j < len(a)-1; j++ {
			if f(a[j], a[j+1]) > 0 {
				return false
			}
		}
		return true
	}
	if f(a[0], a[1]) > 0 {
		for j := 0; j < len(a)-1; j++ {
			if f(a[j], a[j+1]) < 0 {
				return false
			}
		}
	}
	return true
}
