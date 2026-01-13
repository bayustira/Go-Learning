package main

import (
	"fmt"

	"example.com/bank_application/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)
	if err != nil {
		fmt.Println("Error reading balance:", err)
		fmt.Println("-----------------------------------------")
		// panic("Can't Continue, Sorry")
	}

	fmt.Println("Welcome to the Bank Application")
	fmt.Println("Reach us 24/7 at", randomdata.PhoneNumber())

	for {
		presentMenu()

		var choice int
		fmt.Print("Enter your choice (1-4): ")
		fmt.Scan(&choice)

		fmt.Println("You selected option:", choice)

		switch choice {
		case 1:
			fmt.Println("Your current account balance is:", accountBalance)
		case 2:
			var depositAmount float64
			fmt.Print("Enter amount to deposit: ")
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Deposit amount must be positive.")
				continue
			}
			accountBalance += depositAmount
			fmt.Println("Deposit successful. New balance is:", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		case 3:
			var withdrawAmount float64
			fmt.Print("Enter amount to withdraw: ")
			fmt.Scan(&withdrawAmount)
			if withdrawAmount <= 0 {
				fmt.Println("Withdrawal amount must be positive.")
				continue
			}
			if withdrawAmount > accountBalance {
				fmt.Println("Insufficient funds.")
				continue
			}
			accountBalance -= withdrawAmount
			fmt.Println("Withdrawal successful. New balance is:", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		default:
			fmt.Println("Thank you for using the Bank Application. Goodbye!")
			fmt.Println("Exiting the application.")
			return
		}
	}
}
