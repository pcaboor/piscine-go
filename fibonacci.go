package piscine

func Fibonacci(index int) int {
	// index start 0
	if index < 0 {
		return -1
	}
	if index == 0 { // on sait que le 0 est le 1 = 0 et 1
		return 0
	}
	if index == 1 {
		return 1
	}
	return Fibonacci(index-1) + Fibonacci(index-2) // index 0 = 0, index 1 = 1 donc (index0 + index1) = index2 soit 1 + 0 = 1
}
