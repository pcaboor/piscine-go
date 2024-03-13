package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:] // Exclure le nom du programme
	for i := 0; i <= 127; i++ {
		for _, arg := range args {
			if len(arg) > 0 && int(arg[0]) == i {
				fmt.Println(arg)
			}
		}
	}
}
