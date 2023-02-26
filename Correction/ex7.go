package main

import "fmt"

func main() {
	var n int

	fmt.Print("Enter a positive integer: ")
	fmt.Scanln(&n)

	// first two terms of the Fibonacci sequence
	a, b := 0, 1

	// print the Fibonacci sequence up to the nth term
	for i := 1; i <= n; i++ {
		fmt.Printf("%d ", a)

		// compute the next term of the sequence
		a, b = b, a+b
	}
	fmt.Println()
}
