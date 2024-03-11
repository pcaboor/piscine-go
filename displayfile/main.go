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
		fmt.Println(err.Error())
		return
	}
	arr := make([]byte, 1024)
	_, err = file.Read(arr)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Print(string(arr))
}
