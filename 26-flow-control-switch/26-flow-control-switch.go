package main

import "fmt"

func main() {
	foo := 4

	switch {
	case foo < 2:
		fmt.Println("Foo is less than 2")
	case foo == 2:
		fmt.Println("Foo is equal to 2")
	case foo == 3:
		fmt.Println("Foo is equal to 3")
	default:
		fmt.Println("Foo is greater than 3")
	}
}
