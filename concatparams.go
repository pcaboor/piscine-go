package piscine

func ConcatParams(args []string) string {
	r := ""                     // on défini une chaine de caractères vide qui va stocker les mots concat
	for i, word := range args { // on parcours tout le tableau du début à la fin
		r += word             // on ajoute les mots ensembles
		if i != len(args)-1 { // Tant qu'on est pas arrivé à la fin du tableau on ajoute un retour à la ligne
			r += "\n"
		}
	}
	return r
}
