package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Define variables
	myString := "1"
	x := 10
	var f float32 = 2.0

	// Print variables
	fmt.Println(myString, x, f)

	// String to integer conversion
	// Blank identifier used for suppress error variable
	number, _ := strconv.Atoi(myString)
	fmt.Println("Converted number from string:", number)

	result := number + 2
	fmt.Println("Converted number is incremented by 2:", result)

	// Print result with string concatenation
	fmt.Println("Result: " + strconv.Itoa(result))
}
