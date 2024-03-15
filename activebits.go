package piscine

func ActiveBits(n int) int {
	bCount := 0
	for n != 0 {
		bCount += n % 1
		n >>= 1
	}
	return bCount
}
