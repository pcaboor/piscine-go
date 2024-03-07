package piscine

func AppendRange(min, max int) []int {
	var answer []int
	if min >= max {
		return answer
	}
	for i := min; i < max; i++ {
		answer = append(answer, i)
	}
	return answer
}
