package piscine

func UltimateDivMod(a *int, b *int) {
	tempPoint := *a
	*a = *a / (*b)
	*b = tempPoint % (*b)
}
