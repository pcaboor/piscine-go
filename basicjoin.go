package piscine

func BasicJoin(elems []string) string {
	res := ""                    // on défini une chaine de caractères vide qui va stocker les mots concat
	for _, word := range elems { // on parcours tout le tableau du début à la fin
		res += word // on ajoute les mots ensembles
	}
	return res
}
