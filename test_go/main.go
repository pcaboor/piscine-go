package main

import (
	"fmt"
	"piscine"
)

func main() {
	tab1 := []string{"Hello", "how", "are", "you"}
	tab2 := []string{"This", "1", "is", "4", "you"}
	answer1 := piscine.CountIf(piscine.IsNumericNumb2, tab1)
	answer2 := piscine.CountIf(piscine.IsNumericNumb2, tab2)
	fmt.Println(answer1)
	fmt.Println(answer2)
}
