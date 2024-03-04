package piscine

import (
	"fmt"
)

func FishAndChips(nb int) {
	if nb < 0 {
		fmt.Println("error: is negative")
	} else if nb != nb%2 && nb != nb%3 {
		fmt.Println("error: non divisible")
	} else if nb == nb%2 && nb == nb%3 {
		fmt.Println("fish and chips")
	} else if nb%3 == 0 {
		fmt.Println("chips")
	} else if nb%2 == 0 {
		fmt.Println("fish")
	}

}
