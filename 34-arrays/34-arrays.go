package main

import "fmt"

func main() {
	// General array definition
	// <variable_name> := [<array size>]<type>{<... values>}

	// Define first, assign it later...
	intArray := [3]int{}
	intArray[0] = 1
	intArray[1] = 4
	intArray[2] = 5

	// Define and assign same time...
	// If you initialize the array when defining, you do not need to specify the size of array.
	// Use "..." instead of length.
	stringArray := [...]string{"hello", "this", "is", "our", "string", "array"}

	// iterate array
	for _, value := range stringArray {
		fmt.Print(value, " ")
	}

	// iterate array alternative
	for i := 0; i < len(stringArray)-1; i++ {
		fmt.Print(stringArray[i], " ")
	}

	// Print length of array
	fmt.Println("\nLength of string array:", len(stringArray))

	// Print capacity of array
	fmt.Println("Capacity of string array:", cap(stringArray))
}