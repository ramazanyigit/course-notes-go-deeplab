package main

import "fmt"

func main() {
	var message = "Hello Go!"

	// With auto type matching, we can define multi variables in one line with different types.
	var a, b, c = 3, true, 5.5

	fmt.Println(message, a, b, c)
}