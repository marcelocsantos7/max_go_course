package main

import (
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName string
	birthdate string
	age int
	createdAt time.Time
}

func main() {
	user := User {
		firstName: getUserData("Please enter your first name: "),
		lastName: getUserData("Please enter your last name: "),
		birthdate: getUserData("Please enter your birthdate (MM/DD/YYYY): "),
		createdAt: time.Now(),
	}
	outputUserDetails(user)
}

func outputUserDetails(userDetails User) {
	fmt.Println(userDetails.firstName, userDetails.lastName, userDetails.birthdate, userDetails.createdAt)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
