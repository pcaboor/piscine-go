package piscine

func AppendRange(min, max int) []int {
	var answer []int
	if min >= max {
		return nil
	}
	for i := min; i < max; i++ { // i < max == max - 1 soit l'avant dernière valeur du tableau
		answer = append(answer, i)
	}
	return answer
}
