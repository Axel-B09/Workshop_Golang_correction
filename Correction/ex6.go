package main

import "fmt"

func main() {
	var n int

	fmt.Print("Enter a positive integer: ")
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		fmt.Printf("%d ", 2*i)
	}
	fmt.Println()
}
