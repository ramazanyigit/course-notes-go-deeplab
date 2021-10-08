package main

import (
	"fmt"
	"time"
)

func main() {
	///for <variable def>; <condition>; <after action> {
	//	<body>
	//}
	for i := 0; i < 5; i++ {
		fmt.Println("Current i value:", i)
	}

	// Infinite loop examples
	for ;; {
		fmt.Println("infinite loop", time.Now())
	}

	for {
		fmt.Println("infinite loop", time.Now())
	}
}
