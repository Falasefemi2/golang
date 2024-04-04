package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64 
	var years float64
	var expectedReturnRate float64

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	fmt.Print("expectedRate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formatted := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future value (adjusted for inflation): %.1f\n", futureRealValue)

	// fmt.Printf("Future value: %.0f\nFuture Value (adjuested for inflation): %.0f", futureValue, futureRealValue)
	fmt.Print(formatted, formattedRFV)
}