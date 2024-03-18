package piscine

func StringToInt(str string) int {
	if len(str) == 0 {
		return 0
	}
	var nb int
	negative := false
	for i, chr := range str {
		if i == 0 && chr == '-' {
			negative = true
			continue
		}
		digit := int(chr - '0')
		if digit < 0 || digit > 9 {
			return 0
		}
		nb = nb*10 + digit
	}
	if negative {
		nb = -nb
	}
	return nb
}
