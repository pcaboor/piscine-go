package piscine

func Sqrt(nb int) int {

	if nb%2 != 0 {
		return 0
	} else {
		for i := 1; i <= nb; i++ {
			if i*i == nb {
				return i
			} else if i*i > nb {
				return 0
			}
		}
		return 0
	}
}
