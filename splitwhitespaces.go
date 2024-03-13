package piscine

func SplitWhiteSpaces(s string) []string {
	var t []string
	result := -1             // gérer les espaces blanc
	for i, word := range s { // boucle qui parcours toutes les valeurs du tableau de string
		if word == ' ' || word == '\n' || word == '	' { // vérifie les espaces, les tab et les retours à la ligne
			if len(s[result+1:i]) > 0 { // longueur de la sous chaine entre la dernière position de l'espace et la position actuelle
				t = append(t, s[result+1:i])
			}
			result = i
		}
	}
	if result != len(s)-1 { // tant qu'on ai pas à la fin de la chaine de caractère on prend en compte les espaces.
		t = append(t, s[result+1:])
	}
	return t
}
