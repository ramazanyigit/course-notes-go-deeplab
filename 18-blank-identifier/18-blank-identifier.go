package main

import "fmt"

func main() {
	//Normally methods can return data and error at the sametime.
	//If you do not want to handle error, you do not use the err variable.
	//But in go, usage of non-used variables is a compiler error.
	//To avoid this you can use "blank identifier" (_)

	//Example:
	//Use this:
	_, data := method()

	//instead of:
	err, data := method()

	fmt.Println(data)

	// You can use blank identifier on variable assigment
	// But the blank identifier mostly used on function/method returns.
	var _, x, _ = 1000, 9, 0
}
