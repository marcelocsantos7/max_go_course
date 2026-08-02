package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser, err := user.New(userFirstName, userLastName, userBirthdate)

	var name str = "Admin"
	name.log()

	if err != nil {
		fmt.Println(err)
		return
	}

	appUser.OutputUserDetails()

	appAdmin := user.NewAdmin("email@email.com", "23456")
	appAdmin.OutputUserDetails()

	fmt.Println("######################")

	// user.ClearUserName()
	// user.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
