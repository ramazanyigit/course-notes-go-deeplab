package main

import "fmt"

func main() {
	str1 := "The quick red fox"
	str2 := "jumped over"
	str3 := "the lazy brown dog."
	aNumber := 42
	isTrue := true

	stringLength, err := fmt.Println(str1, str2, str3)

	if err == nil {
		fmt.Println("String length:", stringLength)
	}

	fmt.Printf("Value of aNumber: %v\n", aNumber)
	fmt.Printf("Value of isTrue: %v\n", isTrue)
	fmt.Printf("Value of aNumber as float: %.3f\n", float64(aNumber))

	fmt.Printf("Data types: str1-%T, str2-%T, str3-%T, aNumber-%T, and isTrue-%T\n", str1, str2, str3, aNumber, isTrue)

	dataTypesText := fmt.Sprintf("Data types: str1-%T, str2-%T, str3-%T, aNumber-%T, and isTrue-%T\n", str1, str2, str3, aNumber, isTrue)
	fmt.Print(dataTypesText)
}