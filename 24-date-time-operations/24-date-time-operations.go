package main

import (
	"fmt"
	"time"
)

func main() {
	// Creating a specific date
	t := time.Date(2021, time.October, 7, 12, 21, 34, 0, time.UTC)
	fmt.Printf("My launch at %s\n", t)
	fmt.Println("---------")

	// Create a date that contains current date/time information
	now := time.Now()
	fmt.Printf("Current time is %s\n", now)
	fmt.Println("---------")

	// Print day, month, year, hour and minute for "t"
	fmt.Println("Current Date:", t.Day(), t.Month(), t.Year())
	fmt.Printf("Time: %d:%d\n", t.Hour(), t.Minute())
	fmt.Println("---------")

	// Add one day to "t"
	tomorrow := t.AddDate(0, 0, 1)
	fmt.Printf("Tomorrow is %v, %v ,%v, %v \n",
		tomorrow.Weekday(),
		tomorrow.Month(),
		tomorrow.Day(),
		tomorrow.Year())


	longFormat := "Monday, January 2, 2006"
	shortFormat := "1/2/06"

	fmt.Println("Tomorrow is (long)", tomorrow.Format(longFormat))
	fmt.Println("Tomorrow is (short)", tomorrow.Format(shortFormat))
}
