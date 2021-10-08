package main

import "fmt"

// General function definition
// args = args | args, <variable_name> <type>
// return types = return types | return types, <type>
//
// func <function_name>(<args>) <return types> {
// 		<body>
// }
//
// Note: Functions will be passed

// Simple example
func add(x int, y int) int {
	return x+y
}

// Same but args can be defined as this
func addAlternative(x, y int) int {
	return x+y
}

// Multiple results function
func swap(x, y string) (string, string) {
	return y, x
}

// Dynamic args function
func addMulti(terms ...int) (int, int) {
	result := 0
	for _, term := range terms {
		result += term
	}

	return len(terms), result
}

// Named return values function
func addMultiNamed(terms ...int) (termLength, sum int) {
	termLength = len(terms)
	for _, term := range terms {
		sum += term
	}

	return
}

func main() {
	fmt.Println(add(123, 44))
	fmt.Println(addAlternative(123, 44))
	fmt.Println(swap("a", "b"))

	numTerms, sum := addMulti(1, 3, 5, 10, 2222)
	fmt.Println("Added Multi:", numTerms, "terms for a total of", sum)


	numTerms, sum = addMultiNamed(1, 3, 5, 10, 22322)
	fmt.Println("Added Named:", numTerms, "terms for a total of", sum)

	// Example Anonymous Function
	addMultiFunc := func (terms ...int) (termLength, sum int) {
		termLength = len(terms)
		for _, term := range terms {
			sum += term
		}

		return
	}

	numTerms, sum = addMultiFunc(1, 3, 5, 10, 22322)
	fmt.Println("Added Func:", numTerms, "terms for a total of", sum)
}