package piscine

func Capitalize(s string) string {
	if len(s) <= 0 {
		return ""
	}
	temp := []rune(s)
	for i := 0; i < len(temp); i++ {
		if temp[i] >= 'A' && temp[i] <= 'Z' || temp[i] >= 'a' && temp[i] <= 'z' || temp[i] >= '0' && temp[i] <= '9' {
			if temp[i] >= 'a' && temp[i] <= 'z' {
				temp[i] -= 32
			}
			i++
			for i < len(temp) && (temp[i] >= 'a' && temp[i] <= 'z' || temp[i] >= 'A' && temp[i] <= 'Z' || temp[i] >= '0' && temp[i] <= '9') {
				if temp[i] >= 'A' && temp[i] <= 'Z' {
					temp[i] += 32
				}
				i++
			}
		}
	}
	return string(temp)
}