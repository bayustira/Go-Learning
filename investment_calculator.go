package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)
	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)
	outputText("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)
	// futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedFRV := fmt.Sprintf("Future Value (adjusted for Inflation): %.2f\n", futureRealValue)

	//fmt.Printf("Future Value: %.2f\nFuture Value (adjusted for Inflation): %.2f\n", futureValue, futureRealValue)
	//fmt.Println("Future Value:", futureValue)
	//fmt.Println("Future Real Value:", futureRealValue)
	fmt.Print(formattedFV, formattedFRV)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (futureValue float64, futureRealValue float64) {
	futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue = futureValue / math.Pow(1+inflationRate/100, years)
	return futureValue, futureRealValue
}
