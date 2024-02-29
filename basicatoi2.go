package piscine

func BasicAtoi2(s string) int {
	myResult := 0
	for _, numb := range s {

		if numb < '0' || numb > '9' {
			return 0
		}
		myResult = myResult*10 + int(numb-'0')
	}
	return myResult
}
