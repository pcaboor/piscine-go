package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	answer := make([]int, max-min)
	for i := range answer {
		answer[i] = min + i
	}
	return answer
}
