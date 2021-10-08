package main

import "fmt"

func main() {
	// Define primes array
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// Create an integer slice between indices 1 (inclusive) and 4 (exclusive)
	slice := primes[1:4]

	fmt.Println(slice)
	fmt.Println("Length of slice:", len(slice))
	fmt.Println("Capacity of slice:", cap(slice))

	// Slices are points array's elements.
	// It do not contains any data. It POINTS!

	// You can use predefined slice limits.
	// Lower limit 0, upper limit = length of array.
	// The expressions are same:
	// primes[0:6]
	// primes[:6]
	// primes[0:]
	// primes[:]

	// Empty slices are defined as "nil".
	// Length and capacity values will be 0 for an empty slice.
}