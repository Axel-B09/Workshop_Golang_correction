package main

import "fmt"

func main() {
	var num int

	fmt.Print("Enter an integer: ")
	fmt.Scanln(&num)

	if num%2 == 0 {
		fmt.Printf("%d is even\n", num)
	} else {
		fmt.Printf("%d is odd\n", num)
	}
}
