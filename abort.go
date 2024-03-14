package piscine

func Abort(a, b, c, d, e int) int {
	t := []int{a, b, c, d, e}

	for i := 0; i < len(t)-1; i++ {
		for j := 0; j < len(t)-i-1; j++ {
			if t[j] < t[j+1] {
				t[j], t[j+1] = t[j+1], t[j]
			}
		}
	}
	return t[2]
}
