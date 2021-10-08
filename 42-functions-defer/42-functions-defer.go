package main

import "fmt"

func main() {
	// With defer used in a line, the line will be executed, after the function execution is done.
	// Main function will be printed hello world
	defer fmt.Println("world")

	fmt.Println("hello")
}