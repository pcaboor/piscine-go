package piscine

func BasicAtoi(s string) int {
	myResult := 0
	for _, numb := range s {
		myResult = myResult*10 +
			int(numb-'0')
	}
	return myResult

}
