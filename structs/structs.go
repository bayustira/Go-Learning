package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func main() {
	userFirstName := getUserdata("Please enter your first name: ")
	userLastName := getUserdata("Please enter your last name: ")
	userBirthDate := getUserdata("Please enter your birth date (MM/DD/YYYY): ")

	var appUser user
	appUser = user{
		firstName: userFirstName,
		lastName:  userLastName,
		birthDate: userBirthDate,
		createdAt: time.Now(),
	}

	//do something with the collected data
	appUser.outputUserDetails()
}

func (u user) outputUserDetails() {
	fmt.Printf("First Name: %s\nLast Name: %s\nBirth Date: %s\n", u.firstName, u.lastName, u.birthDate)
}

func getUserdata(promptText string) string {
	fmt.Print(promptText)
	var input string
	fmt.Scan(&input)
	return input
}
