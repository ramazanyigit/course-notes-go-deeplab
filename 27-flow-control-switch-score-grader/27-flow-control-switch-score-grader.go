package main

import (
	"fmt"
	"os"
)

func main() {
	var score float64
	var grade byte
	fmt.Print("Enter the score of your exam: ")
	fmt.Scanf("%v", &score)

	switch {
	case score <= 49:
		grade = 'F'
	case score <= 59:
		grade = 'E'
	case score <= 69:
		grade = 'D'
	case score <= 79:
		grade = 'C'
	case score <= 89:
		grade = 'B'
	case score <= 100:
		grade = 'A'
	default:
		fmt.Println("Please enter a score between 0-100")
		os.Exit(1)
	}

	fmt.Printf("Your grade is %c", grade)
}
