package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthdate string
	age       int
	createdAt time.Time
}

func (u *user) outputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthdate, u.createdAt)
}

func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func newUser() *user {
	return &user{
		firstName: getUserData("Please enter your first name: "),
		lastName:  getUserData("Please enter your last name: "),
		birthdate: getUserData("Please enter your birthdate (MM/DD/YYYY): "),
		createdAt: time.Now(),
	}
}

func main() {
	user := newUser()

	user.outputUserDetails()
	user.clearUserName()
	user.outputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
