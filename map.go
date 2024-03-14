package piscine

func Map(f func(int) bool, a []int) []bool {
	array := make([]bool, len(a))
	for index, n := range a {
		array[index] = f(n)
	}
	return array
}
