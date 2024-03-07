package piscine

func ConcatParams(args []string) string {
	answer := ""
	for _, char := range args {
		answer += char + "\n"
	}
	return answer
}
