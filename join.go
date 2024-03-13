package piscine

func Join(strs []string, sep string) string {
	res := "" // on défini une chaine de caractères vide qui va stocker les mots concat
	for i, word := range strs {
		res += word
		if i != len(strs)-1 {
			res += sep
		}
	}
	return res
}
