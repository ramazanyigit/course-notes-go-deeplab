package main

import "fmt"

func main() {
	var i int = 55

	// Cast integer to float64
	var f float64 = float64(i)

	// Cast float64 to unsigned int
	var u uint = uint(f)

	fmt.Println(i, f, u)

	//In Go, all castings must be done explicitly.
	//Casting in general:
	//var destination type = type(source)
}