package piscine

func CountIf(f func(string) bool, tab []string) int {
	count := 0
	for _, n := range tab {
		if f(n) {
			count++
		}
	}
	return count
}

func IsNumericNumb2(s string) bool {
	for _, goodLetter := range s {
		if goodLetter < '0' || goodLetter > '9' {
			return false
		}
	}
	return true
}
