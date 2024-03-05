package piscine

func Index(s string, toFind string) int {
	for index := 0; index <= len(s)-len(toFind); index++ {
		if s[index:index+len(toFind)] == toFind {
			return index
		}
	}
	return -1
}
