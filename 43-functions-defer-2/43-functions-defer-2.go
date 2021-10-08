package main

import "fmt"

var isConnected = false

func main() {
	fmt.Println("Connection status:", isConnected)
	databaseProcessing()
	fmt.Println("Connection status:", isConnected)
}

func databaseProcessing() {
	defer disconnect()

	connect()
	fmt.Println("Deferring disconnect")
	fmt.Printf("Connection open: %v\n", isConnected)
	fmt.Println("Doing something")
}

func connect() {
	isConnected = true
	fmt.Println("Connected to database!")
}

func disconnect() {
	isConnected = false
	fmt.Println("Disconnected!")
}