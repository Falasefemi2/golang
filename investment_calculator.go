package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5


func main() {
	var investmentAmount float64 
	var years float64
	var expectedReturnRate float64

	// fmt.Print("Investment Amount: ")
	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	// fmt.Print("Years: ")
	outputText("Years: ")
	fmt.Scan(&years)

	// fmt.Print("expectedRate: ")
	outputText("expectedRate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue, futureRealValue := calcculateFutureValue(investmentAmount, expectedReturnRate, years)

	// futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formatted := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future value (adjusted for inflation): %.1f\n", futureRealValue)

	// fmt.Printf("Future value: %.0f\nFuture Value (adjuested for inflation): %.0f", futureValue, futureRealValue)
	fmt.Print(formatted, formattedRFV)
}


func outputText(text string) {
	fmt.Print(text)
}

func calcculateFutureValue(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv := fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
	// return
}