package piscine

import "fmt"

func FromTo(from int, to int) string {
	if from < 0 || from > 99 && (to < 0 || to > 99) {
		fmt.Println("Invalid")
	}
	var res string
	step := 1

	if from > to {
		step = -1
	}

	for i := from; i <= to+step; i += step {
		if i < 10 {
			res += "0"
		}
		res += fmt.Sprintf("%0d, ", i)

	}
	return res[:len(res)-2] + "\n"

}
