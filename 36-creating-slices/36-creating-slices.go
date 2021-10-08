package main

import "fmt"

func main() {
	// Define primes array
	primes := [...]int{2, 3, 5, 7, 11, 13}

	// Define a slice and print it's value
	primesSlice := primes[:]
	fmt.Println("Value of slice:", primesSlice)

	// Modify the slice value at index 0
	primesSlice[0] = 19

	// Print array and slice
	fmt.Println("Modified value of array:", primes)
	fmt.Println("Modified value of slice:", primesSlice)

	// ===========================================================
	// Create a string slice.
	colors := []string {"Red", "Green", "Blue"}
	fmt.Println("Colors: ", colors)

	// Add a value to slice.
	colors = append(colors, "Purple")
	fmt.Println("New-Colors:", colors)

	// Remove first value from slice.
	colors = append(colors[1:])
	fmt.Println("Removed-first:", colors)

	// Remove last value from slice.
	colors = append(colors[:len(colors)-1])
	fmt.Println("Removed-last:", colors)
}