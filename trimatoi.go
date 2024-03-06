package piscine

func TrimAtoi(s string) int {
	myResult := 0
	negative := false
	for _, numb := range s {
		if numb == '-' && myResult == 0 {
			negative = true
		}
		if numb >= '0' && numb <= '9' {
			myResult = myResult*10 + int(numb-48)
		}
	}
	if negative {
		myResult = -myResult
	}
	return myResult
}
