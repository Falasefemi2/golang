// package main

// import (
// 	"fmt"
// 	"time"
// )

// type User struct {
// 	firstName string
// 	lastName  string
// 	birthdate string
// 	createdAt time.Time
// }

// func (u User) outPutData() {
// 	fmt.Println(u.firstName, u.lastName, u.birthdate)
// }

// // receiver argument

// func (u *User) clearUserName() {
// 	u.firstName = ""
// 	u.lastName = ""
// }

// func newUser(firstName, lastName, birthdate string) *User {
// 	return &User{
// 		firstName: firstName,
// 		lastName:  lastName,
// 		birthdate: birthdate,
// 		createdAt: time.Now(),
// 	}
// }

// func main() {
// 	firstName := getUserData("Please enter your first name: ")
// 	lastName := getUserData("Please enter your last name: ")
// 	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

// 	var appUser *User

// 	appUser = newUser(firstName, lastName, birthdate)

// 	// ... do something awesome with that gathered data!

// 	// fmt.Println(firstName, lastName, birthdate)

// 	appUser.outPutData()
// 	appUser.clearUserName()
// 	appUser.outPutData()
// }

// func getUserData(promptText string) string {
// 	fmt.Print(promptText)
// 	var value string
// 	fmt.Scan(&value)
// 	return value
// }

package main

import (
	"fmt"
	"strconv"
)

type User struct {
	userName string
	password string
	age      int
}

func main() {
	userName := getUser("Enter username: ")
	password := getUser("Enter password: ")
	age := getUser("Enter age: ")

	ageInt, err := strconv.Atoi(age)
	if err != nil {
		fmt.Println("Invalid age. Please enter a valid number.")
		return
	}

	user := User{
		userName: userName,
		password: password,
		age:      ageInt,
	}

	fmt.Println(user)
}

func getUser(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
