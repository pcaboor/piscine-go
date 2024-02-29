package piscine

func Swap(a *int, b *int) {
	tempPoint := *a
	*a = *b
	*b = tempPoint

}
