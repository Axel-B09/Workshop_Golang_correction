package main

import (
	"fmt"
)

func main() {
	var str string

	fmt.Print("Enter a string: ")
	fmt.Scanln(&str)

	// Convert the string to a slice of runes (Unicode code points)
	runes := []rune(str)

	// Reverse the slice of runes in-place
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// Convert the slice of runes back to a string and print it
	reversedStr := string(runes)
	fmt.Println("Reversed string:", reversedStr)
}
