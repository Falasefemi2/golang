package models

import (
	"errors"

	"example.com/restapi/db"
	"example.com/restapi/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := "INSERT INTO users (email, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)

	if err != nil {
		return nil
	}

	userId, err := result.LastInsertId()

	u.ID = userId
	return err
}

// func (u User) ValidateCredentials() error {
// 	query := "SELECT email, password FROM users WHERE email = ?"
// 	row := db.DB.QueryRow(query, u.Email)

// 	var retrievedPassword string
// 	err := row.Scan(&retrievedPassword)

// 	if err != nil {
// 		return errors.New("password is not valid")
// 	}

// 	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

// 	if !passwordIsValid {
// 		return errors.New("password is not valid")
// 	}

// 	return nil
// }

func (u User) ValidateCredentials() error {
	query := "SELECT email, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrievedEmail, retrievedPassword string
	err := row.Scan(&retrievedEmail, &retrievedPassword)

	if err != nil {
		return errors.New("user not found")
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("password is not valid")
	}

	return nil
}
