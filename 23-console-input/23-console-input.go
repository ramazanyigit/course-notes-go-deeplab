package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Read a string for a text
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	str, _ := reader.ReadString('\n')
	fmt.Println(str)

	// Read a string for number parsing
	fmt.Print("Enter a number: ")
	str, _ = reader.ReadString('\n')

	// Parse string as float.
	f, err := strconv.ParseFloat(strings.TrimSpace(str), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Value of number:", f)
	}
}
