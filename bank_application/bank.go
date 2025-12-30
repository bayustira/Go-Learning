package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func readBalanceFromFile() (float64, error) {
	// Placeholder function to simulate reading balance from a file
	data, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		return 1000, errors.New("Failed to find balance file")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000, errors.New("Failerd to parse stored balance value")
	}
	return balance, nil
}

func writeToBalanceFile(balance float64) {
	// Placeholder function to simulate writing balance to a file
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {
	var accountBalance, err = readBalanceFromFile()
	if err != nil {
		fmt.Println("Error reading balance:", err)
		fmt.Println("-----------------------------------------")
		// panic("Can't Continue, Sorry")
	}

	fmt.Println("Welcome to the Bank Application")
	
	for {
		fmt.Println("What would you like to do today?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit funds")
		fmt.Println("3. Withdraw funds")
		fmt.Println("4. Exit")

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
			writeToBalanceFile(accountBalance)
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
			writeToBalanceFile(accountBalance)
		default:
			fmt.Println("Thank you for using the Bank Application. Goodbye!")
			fmt.Println("Exiting the application.")
			return
		}
	}
}
