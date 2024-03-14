package piscine

func Compact(ptr *[]string) int {
	var t []string
	for _, n := range *ptr {
		if n != "" {
			t = append(t, n)
		}
	}
	*ptr = t
	return len(t)
}
