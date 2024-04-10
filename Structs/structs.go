package main

import (
	"fmt"

	"example.com/structs/user"
)

type str string

func (text str) log() {
	fmt.Println(text)
}

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthRate := getUserData("Please enter your birthRate (MM/DD/YYYY): ")

	var name str = "femi"
	name.log()

	var appUser *user.User

	appUser, err := user.NewUser(firstName, lastName, birthRate)

	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("test@gmail.com", "test124")

	admin.User.OutPutData()
	admin.User.ClearUserName()
	admin.User.OutPutData()

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
