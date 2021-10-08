package main

import "fmt"

func main() {
	//Example 1 - Array
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	//for-each for slice pow
	//i = index
	//v = value
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	//Example 2 - Array
	a := [...]string {"a", "b", "c", "d"}

	//i = index
	//get value using index = a[i]
	for i := range a {
		fmt.Println("Array item", i, "is", a[i])
	}

	//Example 3 - Map
	capitals := map[string]string {
		"Turkey": "Istanbul",
		"France": "Paris",
		"Italy": "Roma",
	}

	for key, value := range capitals {
		fmt.Println("Capital of", key, "is", value)
	}
}
