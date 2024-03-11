package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Hello")
		fmt.Println("Hello")
		fmt.Println("^C")
		return
	}
	for i := 1; i < len(os.Args); i++ {
		arr, err := os.ReadFile(os.Args[i])
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Print(string(arr))
	}
}
