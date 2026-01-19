package main

import "fmt"

func main() {
	firstName := getUserdata("Please enter your first name: ")
	lastName := getUserdata("Please enter your last name: ")
	birthDate := getUserdata("Please enter your birth date (MM/DD/YYYY): ")

	//do something with the collected data
	fmt.Printf("First Name: %s\nLast Name: %s\nBirth Date: %s\n", firstName, lastName, birthDate)
}

func getUserdata(promptText string) string {
	var input string
	fmt.Print(promptText)
	fmt.Scanln(&input)
	return input
}
