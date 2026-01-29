package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	userFirstName := getUserdata("Please enter your first name: ")
	userLastName := getUserdata("Please enter your last name: ")
	userBirthDate := getUserdata("Please enter your birth date (MM/DD/YYYY): ")

	var appUser *user.User
	appUser, err := user.New(userFirstName, userLastName, userBirthDate)

	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}

	admin := user.NewAdmin("admin@example.com", "password")

	admin.OutputUserDetails()
	admin.ClearUserDetails()
	admin.OutputUserDetails()

	//do something with the collected data
	appUser.OutputUserDetails()
	appUser.ClearUserDetails()
	appUser.OutputUserDetails()
}

func getUserdata(promptText string) string {
	fmt.Print(promptText)
	var input string
	fmt.Scanln(&input)
	return input
}
