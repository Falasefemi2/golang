package main

import (
	"fmt"
)




func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	// fmt.Print("Revenue: ")
	calculate("Revenue: ")
	fmt.Scan(&revenue)

	// fmt.Print("Expenses: ")
	calculate("Expenses: ")
	fmt.Scan(&expenses)

	// fmt.Print("Tax Rate: ")
	calculate("Tax Rate: ")
	fmt.Scan(&taxRate)

	 ebt, profit, ratio := calculateAmout(revenue, expenses, taxRate)
	 fmt.Printf("%0.1f\n", ebt)
	 fmt.Printf("%0.1f\n", profit)
	 fmt.Printf("%0.3f\n", ratio)

}

func calculate(n string) {
	fmt.Print(n)
}

func calculateAmout(r,e,t float64) (float64, float64, float64) {
	ebt := r - e
	p := ebt * (1-t/100)
	ra := ebt/p
	return ebt, p, ra
}


