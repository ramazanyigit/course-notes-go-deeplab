package main

import "fmt"

// Test it. Short definition variable will be throw error on package scope.
// You need to use var keyword definition.
// Wrong:
// num := 24
// Correct:
// var num = 24

func main() {
	// Short definitions can only use in function bodies.
	// Cannot use on package scope.

	u := 40
	v, n := "abc", true
	message := "Hello Go!"
	a, b, c := 4, false, 24.13

	fmt.Println(u, v, n)
	fmt.Println(message, a, b, c)

	// One-line string
	d := "Text"
	// Character
	e := 'C'
	// Multiline string
	f := `Multiline
	String`

	fmt.Println(d, e)
	fmt.Println(f)

	// Specific type
	var myFloat32 float32 = 44.

	// Complex number
	myComplex := complex(3, 4)

	fmt.Println(myFloat32)
	fmt.Println(myComplex)
}
