package piscine

func SplitWhiteSpaces(s string) []string {
	var tbstr []string
	result := -1 // gérer les espaces
	for i, word := range s {
		if word == ' ' || word == '\n' || word == '	' {
			if len(s[result+1:i]) > 0 {
				tbstr = append(tbstr, s[result+1:i])
			}
			result = i
		}
	}
	if result != len(s)-1 { // tant qu'on ai pas à la fin de la chaine de caractère on prend en compte les espaces.
		tbstr = append(tbstr, s[result+1:])
	}
	return tbstr
}
