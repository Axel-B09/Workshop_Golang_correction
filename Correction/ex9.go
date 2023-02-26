package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Open the file
	file, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Initialize a variable to keep track of the paragraph
	paragraph := ""

	// Read the file line by line
	for scanner.Scan() {
		// If the line is empty, we've reached the end of a paragraph
		if scanner.Text() == "" {
			// Print the paragraph and exit the loop
			fmt.Println(paragraph)
			break
		} else {
			// Append the line to the paragraph
			paragraph += scanner.Text() + "\n"
		}
	}

	// Handle any errors that occurred while reading the file
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
