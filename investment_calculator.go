package main

import (
	"fmt"
)

func main() {
	// const inflationRate = 2.5
	// var investmentAmount float64 
	// var years float64
	// var expectedReturnRate float64

	// fmt.Print("Investment Amount: ")
	// fmt.Scan(&investmentAmount)

	// fmt.Print("Years: ")
	// fmt.Scan(&years)

	// fmt.Print("expectedRate: ")
	// fmt.Scan(&expectedReturnRate)

	// futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	// fmt.Println("future value: ",futureValue)
	// fmt.Println("futureRealValues: ",futureRealValue)

	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	EBT := revenue - expenses
	profit := EBT * (1-taxRate/100)
	ratio := EBT / profit

	outPut := EBT + profit + ratio

	fmt.Println("output: ", outPut)

}