package main

import (
	"fmt"
)

func main() {
	fmt.Println("start")
	// After this line program will be jumped to line 15 and continue the execution.
	// Println on line 14 will be never printed.
	goto FINISH

	fmt.Println("this line will not be printed.")

FINISH:
	fmt.Println("finish!")
}
