package piscine

func Unmatch(a []int) int {
	t := make(map[int]int)
	for _, n := range a {
		t[n]++
	}
	t2 := []int{}
	for n, count := range t {
		if count%2 != 0 {
			t2 = append(t2, n)
		}
	}
	if len(t2) == 1 {
		return t2[0]
	}
	return -1
}
