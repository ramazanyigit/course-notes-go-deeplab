package main

import "fmt"

func main() {
	a := 10
	b := 30
	fmt.Println("A:", a)
	fmt.Println("B:", b)

	sum := a + b
	fmt.Println("Addition operation A+B:", sum)

	difference := a - b
	fmt.Println("Subtraction operation A-B:", difference)

	product :=  a * b
	fmt.Println("Multiply operation A*B:", product)

	division := float64(a) / float64(b)
	fmt.Println("Divide operation A/B:", division)

	remain :=  a % b
	fmt.Println("Modulus operation A % B:", remain)

	a++
	fmt.Println("A is incremented by 1:", a)

	b--
	fmt.Println("A is decremented by 1:", b)
}
