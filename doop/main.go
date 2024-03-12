package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 4 { // vérifie qu'on entre les 3 arguments : value, operator et value2
		return
	}
	x, err := Atoi2(os.Args[1]) // convertir la première valeur en entier int et gestion des erreurs
	if err != nil {
		return
	}
	sign := os.Args[2]          // stock la deuxième valeur entant qu'opérateur et gestion des erreurs
	y, err := Atoi2(os.Args[3]) // convertir la troisième valeur en entier int et gestion des erreurs
	if err != nil {
		return
	}
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
			fmt.Println("No division by 0")
			return
		}
		r = x / y
	case "%":
		if y == 0 {
			fmt.Println("No modulo by 0")
			return
		}
		r = x % y
	default:
		return
	}

	fmt.Println(r)

}

func Atoi2(s string) (int, error) {
	myResult := 0
	ponc := 1

	for a, b := range s {
		if a == 0 && (b == '+' || b == '-') {
			if b == '-' {
				ponc = -1
			}
		} else if b < '0' || b > '9' {
			return 0, fmt.Errorf("Error")
		} else {
			myResult = myResult*10 + int(b-'0')
		}
	}
	return myResult * ponc, nil
}
