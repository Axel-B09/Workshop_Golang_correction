package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	filename := "example.txt"

	// Read the entire file into memory
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	// Print the file contents as a string
	fmt.Println(string(data))
}
