package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

type Admin struct {
	User
	email    string
	password string
}

func (u User) OutputUserDetails() {
	fmt.Printf("First Name: %s\nLast Name: %s\nBirth Date: %s\n", u.firstName, u.lastName, u.birthDate)
}

func (u *User) ClearUserDetails() {
	u.firstName = ""
	u.lastName = ""
	u.birthDate = ""
}

func NewAdmin(email, password string) Admin {
	return Admin{
		User: User{
			firstName: "Admin",
			lastName:  "User",
			birthDate: "01/01/2000",
			createdAt: time.Now(),
		},
		email:    email,
		password: password,
	}
}

func New(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("all fields are required")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}
