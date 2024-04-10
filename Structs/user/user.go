package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthRate string
	createdAt time.Time
}

// struct embedding
type Admin struct {
	email    string
	password string
	User     User
}

func (u *User) OutPutData() {
	fmt.Println(u.firstName, u.lastName, u.birthRate)
}

// receiver argument

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName:  "ADMIN",
			birthRate: "____",
			createdAt: time.Now(),
		},
	}
}

func NewUser(firstName, lastName, birthRate string) (*User, error) {
	if firstName == "" || lastName == "" || birthRate == "" {
		return nil, errors.New("first name, last name and birthrate are required")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthRate: birthRate,
		createdAt: time.Now(),
	}, nil
}
