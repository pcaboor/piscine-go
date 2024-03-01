package piscine

func Atoi(s string) int {
	myResult := 0
	ponc := 1

	for a, b := range s {
		if a == 0 && (b == '+' || b == '-') {
			if b == '-' {
				ponc = -1
			}
		} else if b < '0' || b > '9' {
			return 0
		} else {

			myResult = myResult*10 + int(b-'0')
		}

	}
	return myResult * ponc
}
