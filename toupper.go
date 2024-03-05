package piscine

func ToUpper(s string) string {
	temp := []rune(s)
	for i := 0; i < len(temp); i++ {
		if temp[i] >= 'a' && temp[i] <= 'z' {
			temp[i] -= 32
		}
	}
	return string(temp)
}
