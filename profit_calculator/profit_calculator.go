package main

import (
	"errors"
	"fmt"
	"os"
)

//Goals
// 1. Validate user input
//   => Show error message & exit if invalid input
// 	 - No Negative numbers
// 	 - No zero values
// 2. Store calculated results in a file
func main() {
	revenue, err := getUserInput("Revenue:")
	if err != nil {
		fmt.Println(err)
		return
	}
	expenses, err := getUserInput("Expenses:")
	if err != nil {
		fmt.Println(err)
		return
	}
	taxRate, err := getUserInput("Tax Rate:")
	if err != nil {
		fmt.Println(err)
		return
	}

	earningBeforeTax, earningAfterTax, ratio := calculateProfit(revenue, expenses, taxRate)

	fmt.Printf("EBT: %.2f\n", earningBeforeTax)  
	fmt.Printf("Profit: %.2f\n", earningAfterTax)
	fmt.Printf("Ratio: %.2f\n", ratio)

	storeReslultsToFile(earningBeforeTax, earningAfterTax, ratio)
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		return 0, errors.New("Invalid input, please enter a positive value")
	}
	return userInput, nil
}

func calculateProfit(revenue, expenses, taxRate float64) (float64, float64, float64) {
	earningBeforeTax := revenue - expenses
	earningAfterTax := earningBeforeTax * (1 - taxRate/100)
	ratio := earningBeforeTax / earningAfterTax
	return earningBeforeTax, earningAfterTax, ratio
}

func storeReslultsToFile(earningBeforeTax, earningAfterTax, ratio float64) {
	results := fmt.Sprintf("EBT: %.2f\nProfit: %.2f\nRatio: %.2f\n", earningBeforeTax, earningAfterTax, ratio)
	err := os.WriteFile("profit_results.txt", []byte(results), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}
