package piscine

func SplitWhiteSpaces(s string) []string {
	var tbstr []string
	result := -1
	for i, letter := range s {
		if letter == ' ' || letter == '\n' || letter == '\t' {
			if len(s[result+1:i]) > 0 {
				if s[result+1:i] != "" {
					tbstr = append(tbstr, s[result+1:i])
				}
				result = i
			}
		}
	}
	if result != len(s)-1 {
		tbstr = append(tbstr, s[result+1:])
	}
	return tbstr
}
