package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthRate := getUserData("Please enter your birthRate (MM/DD/YYYY): ")

	var appUser *user.User

	appUser, err := user.NewUser(firstName, lastName, birthRate)

	if err != nil {
		fmt.Println(err)
		return
	}

	// ... do something awesome with that gathered data!

	// fmt.Println(firstName, lastName, birthRate)

	appUser.OutPutData()
	appUser.ClearUserName()
	appUser.OutPutData()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
