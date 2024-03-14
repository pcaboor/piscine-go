package main

import (
	"fmt"
	"os"
)

func main() {
	a := os.Args[1:]

	for _, ch := range a {
		if ch == "01" || ch == "galaxy" || ch == "galaxy 01" {
			fmt.Println("Alert!!!")
			return
		}
	}
}
