package main

import "fmt"

func main() {
	a := true
	b := false
	fmt.Println("A:", a)
	fmt.Println("B:", b)

	andOperation := a && b
	fmt.Println("A && B:", andOperation)

	orOperation := a || b
	fmt.Println("A || B:", orOperation)

	notOperation := !a
	fmt.Println("!A:", notOperation)
}
