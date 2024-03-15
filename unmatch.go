package piscine

func Unmatch(a []int) int {
	t := make(map[int]int)
	for _, n := range a {
		t[n]++
	}
	for n, count := range t {
		if count%2 != 0 {
			return n
		}
	}
	return -1
}
