package main

import (
	"fmt"
)

func main() {
	fmt.Println("First user creation")

	firstUser := &User{
		ID:        1,
		FirstName: "Ramazan",
		LastName:  "Yigit",
		UserName:  "ramazanyigit",
		Age:       23,
		Pay: &Payment{
			Salary: 3555,
			Bonus:  850,
		},
	}

	fmt.Println(firstUser.GetFullName())
	fmt.Println(firstUser.GetPayment())

	// Second way to user creation
	fmt.Println("Second user creation")

	secondUser := NewUser()
	secondUser.FirstName = "Ramazan"
	secondUser.LastName = "Yigit"
	secondUser.Age = 30
	secondUser.UserName = "ramazanyigit"
	secondUser.Pay.Salary = 1000
	secondUser.Pay.Bonus = 200

	fmt.Println(secondUser.GetFullName())
	fmt.Println(secondUser.GetPayment())
}

type User struct {
	ID        int
	FirstName string
	LastName  string
	UserName  string
	Age       int
	Pay       *Payment
}

func NewUser() *User {
	u := new(User)
	u.Pay = NewPayment()
	return u
}

type Payment struct {
	Salary float64
	Bonus  float64
}

func NewPayment() *Payment {
	p := new(Payment)
	return p
}

// User Methods

func (u *User) GetFullName() string {
	return u.FirstName + " " + u.LastName
}

func (u *User) GetUserName() string {
	return u.UserName
}

func (u *User) GetPayment() float64 {
	return u.Pay.Salary + u.Pay.Bonus
}
