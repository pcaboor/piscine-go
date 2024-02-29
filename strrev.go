package piscine

func StrRev(s string) string {
	reverse := []rune(s)
	for x, y := 0, len(s)-1; x < y; x, y = x+1, x-1 {
		reverse[x], reverse[y] = reverse[y], reverse[x]
	}
	return string(reverse)
}
