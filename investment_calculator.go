package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)
	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedFRV := fmt.Sprintf("Future Value (adjusted for Inflation): %.2f\n", futureRealValue)

	//fmt.Printf("Future Value: %.2f\nFuture Value (adjusted for Inflation): %.2f\n", futureValue, futureRealValue)
	//fmt.Println("Future Value:", futureValue)
	//fmt.Println("Future Real Value:", futureRealValue)
	fmt.Print(formattedFV, formattedFRV)
}
