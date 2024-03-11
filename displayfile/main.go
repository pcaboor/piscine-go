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
	file, err := os.Open("sample.txt")

	if err != nil {
		fmt.Printf("File name missing")
	}
	arr := make([]byte, 12)
	count, err := file.Read(arr)

	fmt.Println(file.Stat())

	if err != nil {
		fmt.Printf("Error")
		return
	}
	fmt.Printf("%s", arr[:count])
}
