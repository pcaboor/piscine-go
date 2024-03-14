package piscine

func Rot14(s string) string {
	temp := []rune(s)
	for i := 0; i < len(temp); i++ {
		if temp[i] >= 'A' && temp[i] <= 'z' {
			temp[i] += 14
			if (temp[i] > 'Z' && temp[i] <= 'Z'+14) || temp[i] > 'z' {
				temp[i] -= 26
			}
		}
	}
	return string(temp)
}
