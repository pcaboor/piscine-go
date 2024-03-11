package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("File name missing")
		return
	}
	if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
		return
	}
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("File name missing")
		return
	}
	arr := make([]byte, 12)
	_, err = file.Read(arr)
	if err != nil {
		fmt.Printf("File name missing")
		return
	}
	fmt.Println(string(arr))
}
