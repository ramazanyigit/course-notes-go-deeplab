package main

import (
	"fmt"
	"os"
)

func main() {
	userName := os.Getenv("USERNAME")
	userDomain := os.Getenv("USERDOMAIN")
	processorArchitecture := os.Getenv("PROCESSOR_ARCHITECTURE")
	goRoot := os.Getenv("GOROOT")
	goPath := os.Getenv("GOPATH")
	homePath := os.Getenv("HOMEPATH")

	fmt.Println("User Name: " + userName)
	fmt.Println("User Domain Name: " + userDomain)
	fmt.Println("Processor Architecture: " + processorArchitecture)
	fmt.Println("Home Path: " + homePath)
	fmt.Println("Go Path: " + goPath)
	fmt.Println("Go Root Path: " + goRoot)
}
