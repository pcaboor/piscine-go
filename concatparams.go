package piscine

func ConcatParams(args []string) string {
	r := ""
	for i, letter := range args {
		r += letter
		if i != len(args)-1 {
			r += "\n"
		}
	}
	return r
}
