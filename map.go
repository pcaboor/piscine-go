package piscine

func Map(f func(int) bool, a []int) []bool {
	array := make([]bool, len(a))
	for index, n := range a {
		array[index] = f(n)
	}
	return array
}
func Prime(nb int) bool {
	if nb <= 1 {
		return false
	}
	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
