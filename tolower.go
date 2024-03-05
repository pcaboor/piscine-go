package piscine

func ToLower(s string) string {
	temp := []rune(s)
	for i := 0; i < len(temp); i++ {
		if temp[i] >= 'A' && temp[i] <= 'Z' {
			temp[i] += 32
		}
	}
	return string(temp)
}
