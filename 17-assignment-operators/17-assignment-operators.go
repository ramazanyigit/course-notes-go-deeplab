package main

import "fmt"

func main() {
	a := 12
	b := 40

	total := a+b
	fmt.Println("Total:", total)

	total -= 5
	fmt.Println("Decrement total by 5:", total)

	total *= 10
	fmt.Println("Multiply total by 10:", total)

	total /= 20
	fmt.Println("Divide total by 20:", total)

	total %= 5
	fmt.Println("Modulus operation on total by 5:", total)

	fmt.Println("Address of total: ", &total)
}
