package main

import (
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"strings"
)

// Every Go program is made up of packages.
// Programs start running in package main.
// Packages offer clear code separation and allow efficient management of dependencies.

// Example usages from several packages.
func main() {
	// fmt
	fmt.Println("This is a package import example")

	// fmt and math/rand
	fmt.Println("Random number from math/rand package:", rand.Intn(10))

	// strings
	fmt.Println("String contains \"es\":", strings.Contains("test", "es"))
	fmt.Println("Counts of \"t\" in \"test\":", strings.Count("test", "t"))

	// external color package usage
	color.Red("Error Message")
}
