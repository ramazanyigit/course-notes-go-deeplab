package main

import "fmt"

// Create a type that will be used for enum.
type Brand string

// Define enums as constants with created custom type.
const (
	FACEBOOK Brand = "Facebook"
	MICROSOFT Brand = "Microsoft"
	GOOGLE Brand = "Google"
	AMAZON Brand = "Amazon"
	KRON Brand = "Kron"
	IRONSPHERE Brand = "Ironsphere"
)

// A print function for brand.
func PrintBrand(brand Brand) {
	fmt.Println(brand)
}

func main() {
	// Usage examples.
	PrintBrand(GOOGLE)
	PrintBrand(KRON)
	PrintBrand(IRONSPHERE)

	//compile error: undefined: VERIZON
	//PrintBrand(VERIZON)
}
