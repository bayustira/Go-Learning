package main

import "fmt"

//Goals
// 1. Validate user input
//   => Show error message & exit if invalid input
// 	 - No Negative numbers
// 	 - No zero values
// 2. Store calculated results in a file
func main() {
	revenue := getUserInput("Revenue:")
	expenses := getUserInput("Expenses:")
	taxRate := getUserInput("Tax Rate:")

	earningBeforeTax, earningAfterTax, ratio := calculateProfit(revenue, expenses, taxRate)

	fmt.Printf("EBT: %.2f\n", earningBeforeTax)  
	fmt.Printf("Profit: %.2f\n", earningAfterTax)
	fmt.Printf("Ratio: %.2f\n", ratio)
}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}

func calculateProfit(revenue, expenses, taxRate float64) (float64, float64, float64) {
	earningBeforeTax := revenue - expenses
	earningAfterTax := earningBeforeTax * (1 - taxRate/100)
	ratio := earningBeforeTax / earningAfterTax
	return earningBeforeTax, earningAfterTax, ratio
}
