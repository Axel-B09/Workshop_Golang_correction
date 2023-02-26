package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Open the file
	file, err := os.Open("words.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Initialize a variable to keep track of the longest word(s)
	longest := []string{}

	// Read the file line by line
	for scanner.Scan() {
		word := scanner.Text()
		if len(word) > len(longest[0]) {
			// We've found a new longest word
			longest = []string{word}
		} else if len(word) == len(longest[0]) {
			// We've found a tie for the longest word
			longest = append(longest, word)
		}
	}

	// Print the longest word(s)
	fmt.Println("Longest word(s):", longest)
}
