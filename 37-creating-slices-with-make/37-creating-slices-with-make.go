package main

import "fmt"

func main() {
	// Create a slice with make
	numbers := make([]int, 5, 5)
	numbers[0] = 22
	numbers[1] = 233
	numbers[2] = 2123
	numbers[3] = 685
	numbers[4] = 109

	fmt.Println("Numbers:", numbers)
	fmt.Println("Length:", len(numbers))
	fmt.Println("Capacity:", cap(numbers))
	fmt.Println()

	// If we add the sixth element, append will allocate new array with capacity 10 instead of 5 and add the value into it.
	numbers = append(numbers, 111)
	fmt.Println("Changed numbers:", numbers)
	fmt.Println("Changed Length:", len(numbers))
	fmt.Println("Changed Capacity:", cap(numbers))
}