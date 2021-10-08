package main

import (
	"fmt"
)

func main() {
	// Break example
	i := 0
	for {
		if i >= 3 {
			break
		}

		fmt.Println("Current i value:", i)
		i++
	}

	for y := 0; y < 7; y++ {
		if y%2 == 0 { // Do not print even numbers.
			// Stop here and go to next loop part.
			continue
		}

		fmt.Println("Value:", y)
	}
}
