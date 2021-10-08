package main

import "fmt"

func main() {
	a := 10
	b := 10

	if b > a {
		fmt.Println("B is greater than A")
	} else if b == a {
		fmt.Println("B is equal to A")
	} else {
		fmt.Println("B is less than A")
	}

	foo := 1
	if foo == 1 {
		fmt.Println("Foo is equal to 1: Bar")
	}

	// IF WITH VARIABLE DECLARATION
	// "bar" can be used only in if's scope.
	if bar := 2; bar == 1 {
		fmt.Println("Bar is equal to 1: foo")
	} else {
		fmt.Println("Bar is not equal to 1: buz")
	}
}
