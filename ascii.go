package piscine

func Ascii(str string) []byte {
	var ascii []byte
	for _, char := range str {
		ascii = append(ascii, byte(char))
	}
	return ascii
}
