package main

import "fmt"

type Human struct {
	FirstName string
	LastName  string
	Age       int
}

// Default and empty constructor example
func NewHuman() *Human {
	h := new(Human)
	return h
}

// All args constructor
func NewHumanWithArgs(firstName, lastName string, age int) *Human {
	h := new(Human)
	h.FirstName = firstName
	h.LastName = lastName
	h.Age = age

	return h
}

func PrintHuman(human *Human) {
	fmt.Println("Name:", human.FirstName, human.LastName, "- Age:", human.Age)
}

func main() {
	// First way to create instance
	x := &Human{FirstName: "Ramazan", LastName: "Yigit"}
	PrintHuman(x)

	// Second way to create instance
	y := new(Human)
	y.FirstName = "John"
	y.LastName = "Doe"
	PrintHuman(y)

	// Create instance from constructor
	z := NewHuman()
	z.Age = 30
	PrintHuman(z)

	t := NewHumanWithArgs("Ramazan", "Yigit", 80)
	PrintHuman(t)
}
