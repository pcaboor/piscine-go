package piscine

func IntToString(nb int) string {
	if nb == 0 {
		return "0"
	}
	var r string
	negative := false
	if nb < 0 {
		negative = true
		nb -= nb
	}
	for nb != 0 {
		digit := nb % 10
		r = string('0'+byte(digit)) + r
		nb /= 10
	}
	if negative {
		r = "-" + r
	}
	return r
}
