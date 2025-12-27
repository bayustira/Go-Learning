package main

import "fmt"

func main() {
	var accountBalance float64 = 1000.00

	fmt.Println("Welcome to the Bank Application")
	fmt.Println("What would you like to do today?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit funds")
	fmt.Println("3. Withdraw funds")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Enter your choice (1-4): ")
	fmt.Scan(&choice)

	fmt.Println("You selected option:", choice)
	
	if choice == 1 {
		fmt.Println("Your current account balance is:", accountBalance)
	} else if choice == 2 {
		var depositAmount float64
		fmt.Print("Enter amount to deposit: ")
		fmt.Scan(&depositAmount)
		accountBalance += depositAmount
		fmt.Println("Deposit successful. New balance is:", accountBalance)
	} else if choice == 3 {
		var withdrawAmount float64
		fmt.Print("Enter amount to withdraw: ")
		fmt.Scan(&withdrawAmount)
		if withdrawAmount > accountBalance {
			fmt.Println("Insufficient funds.")
		} else {
			accountBalance -= withdrawAmount
			fmt.Println("Withdrawal successful. New balance is:", accountBalance)
		}
	} else {
		fmt.Println("Thank you for using the Bank Application. Goodbye!")
		
	}
}
