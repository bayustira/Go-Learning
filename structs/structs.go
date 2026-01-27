package main

import (
	"errors"
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func (u user) outputUserDetails() {
	fmt.Printf("First Name: %s\nLast Name: %s\nBirth Date: %s\n", u.firstName, u.lastName, u.birthDate)
}

func (u *user) clearUserDetails() {
	u.firstName = ""
	u.lastName = ""
	u.birthDate = ""
}

func newUser(firstName, lastName, birthDate string) (*user, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("all fields are required")
	}

	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

func main() {
	userFirstName := getUserdata("Please enter your first name: ")
	userLastName := getUserdata("Please enter your last name: ")
	userBirthDate := getUserdata("Please enter your birth date (MM/DD/YYYY): ")

	var appUser *user
	appUser, err := newUser(userFirstName, userLastName, userBirthDate)

	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}

	//do something with the collected data
	appUser.outputUserDetails()
	appUser.clearUserDetails()
	appUser.outputUserDetails()
}

func getUserdata(promptText string) string {
	fmt.Print(promptText)
	var input string
	fmt.Scanln(&input)
	return input
}
