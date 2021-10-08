package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {
	// Handling an error
	file, err := os.Open("file.txt")
	checkError(err)

	fmt.Println("File name: ", file.Name())

	// Creating a custom error
	customError := errors.New("This a custom error defined by user")
	fmt.Println("Custom Error:", customError.Error())

	// Creating custom error with fmt
	_ = fmt.Errorf("We cannot read the file: %s", file.Name())
}

// Custom error handler to prevent duplicating code.
func checkError(err error) {
	if err == nil {
		return
	}

	log.Fatalln("ERROR:", err)
}