package main

import (
	"os"
)

func main() {
	if len(os.Args) != 4 { // vérifie qu'on entre les 3 arguments : value, operator et value2
		return
	}
	x := Atoi2(os.Args[1]) // convertir la première valeur en entier int
	sign := os.Args[2]     // stock la deuxième valeur entant qu'opérateur
	y := Atoi2(os.Args[3]) // convertir la troisième valeur en entier int

	var r int

	switch sign {
	case "+":
		r = x + y
	case "-":
		r = x - y
	case "*":
		r = x * y

	case "/":
		if y == 0 {
			os.Stderr.WriteString("No divisible by 0\n")
			return
		}
		r = x / y
	case "%":
		if y == 0 {
			os.Stderr.WriteString("No modulo by 0\n")
			return
		}
		r = x % y
	default:
		return
	}
	os.Stdout.WriteString(its(r) + "\n")
}

func Atoi2(s string) int {
	myResult := 0
	ponc := 1

	for a, b := range s {
		if a == 0 && (b == '+' || b == '-') {
			if b == '-' {
				ponc = -1
			}
		} else if b < '0' || b > '9' {
			os.Exit(0)
		} else {
			myResult = myResult*10 + int(b-'0')
		}
	}
	return myResult * ponc
}

func its(i int) string {
	if i == 0 {
		return "0"
	}
	var r string
	sign := ""

	if i < 0 {
		sign = "-"
		i = -i
	}
	for i > 0 {
		numb := i % 10
		r = string(byte(numb)+'0') + r
		i /= 10
	}
	return sign + r
}
