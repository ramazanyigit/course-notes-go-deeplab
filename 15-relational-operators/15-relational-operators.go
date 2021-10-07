package main

import "fmt"

func main() {
	a := 10
	b := 30
	fmt.Println("A:", a)
	fmt.Println("B:", b)

	isEqual := a == b
	fmt.Println("A is equal to B:", isEqual)

	isNotEqual := a != b
	fmt.Println("A is not equal to B:", isNotEqual)

	isGreater := a > b
	fmt.Println("A is greater than B:", isGreater)

	isLower := a < b
	fmt.Println("A is lower than B:", isLower)

	isGreaterOrEqual := a >= b
	fmt.Println("A is greater than or equal to B:", isGreaterOrEqual)

	isLowerOrEqual := a <= b
	fmt.Println("A is lower than or equal to B:", isLowerOrEqual)
}
