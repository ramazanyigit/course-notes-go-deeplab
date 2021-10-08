package main

import (
	"fmt"
)

func main() {
	// Create while loop using "for"
	sum := 1

	for sum < 10 {
		sum += sum
		fmt.Println("Sum:", sum)
	}
}
