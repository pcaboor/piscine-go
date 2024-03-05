package piscine

func IsPrintable(s string) bool {
	for _, imp := range s {
		if imp == '\n' {
			return false
		}
	}
	return true
}
