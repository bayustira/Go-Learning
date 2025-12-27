package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Bank Application")
	fmt.Println("What would you like to do today?")
	fmt.Println("1. Check bnalance")
	fmt.Println("2. Deposit funds")
	fmt.Println("3. Withdraw funds")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Enter your choice (1-4): ")
	fmt.Scan(&choice)

	fmt.Println("You selected option:", choice)
	// Additional bank application logic would go here
}
