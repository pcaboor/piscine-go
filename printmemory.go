package piscine

import (
	"fmt"
)

func PrintMemory(arr [10]byte) {
	// Display memory as hexadecimal representation
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%02x ", arr[i])
		if i == 3 || i == 7 {
			fmt.Println()
		}
	}
	fmt.Println()

	// Display ASCII graphic characters
	for i := 0; i < len(arr); i++ {
		if arr[i] >= 32 && arr[i] <= 126 {
			fmt.Printf("%c", arr[i])
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println()
}
