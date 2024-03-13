package piscine

func Join(strs []string, sep string) string {
	res := ""                   // on défini une chaine de caractères vide qui va stocker les mots concat
	for i, word := range strs { // on parcours tout le tableau du début à la fin
		res += word
		if i != len(strs)-1 {
			res += ":"
		}
	}
	return res
}
