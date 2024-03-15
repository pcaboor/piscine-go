package piscine

func Unmatch(a []int) int {
	for _, n := range a {
		t := 0
		for _, s := range a {
			if s == n {
				t++
			}
		}
		if t == 1 || t%2 == 1 {
			return n
		}
	}
	return -1
}
