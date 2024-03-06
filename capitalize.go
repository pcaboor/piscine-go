package piscine

func Capitalize(s string) string {
	temp := []rune(s)
	for i := 0; i < len(temp); i++ {
		if temp[0] >= 'a' && temp[0] <= 'z' {
			temp[0] -= 32
		}
	}
	return string(temp)
}
