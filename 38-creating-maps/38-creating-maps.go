package main

import "fmt"

func main() {
	// Create an integer-string map
	firstMap := make(map[int]string)
	fmt.Println(firstMap)
	firstMap[42] = "Foo"
	firstMap[24] = "Bar"
	fmt.Println(firstMap)

	// Create a string-string map
	states := make(map[string]string)
	states["IST"] = "Istanbul"
	states["IZM"] = "Izmir"
	states["ANK"] = "Ankara"
	fmt.Println(states)

	izmir := states["IZM"]
	fmt.Println("IZM:", izmir)

	// Deleting value from map by key
	delete(states, "IZM")

	fmt.Println(states)

	// Adding value to map
	states["ANT"] = "Antalya"

	for k, v := range states {
		fmt.Println(k, v)
	}

	// Get keys of states map
	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	fmt.Println(keys)

	// Print states from keys
	for i := range keys {
		fmt.Println(states[keys[i]])
	}
}