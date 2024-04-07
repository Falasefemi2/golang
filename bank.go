package main

import "fmt"

func main() {
	var accountBalance float64 = 1000


	fmt.Println("Welcome to Go Bank!")
	fmt.Println("What do you want to do")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")

	var choice int 
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	// wantsCheckBalance := choice == 1

	if choice  == 1 {
		fmt.Println("Your bslsnce is", accountBalance)
	} else if choice == 2 {
		fmt.Print("Your deposit: ")
		var depositAmount float64
		fmt.Scan(&depositAmount)

		if depositAmount <= 0 {
			fmt.Println("Invalid amount. Must be greater than zero")
			return
		}

		accountBalance += depositAmount
		fmt.Println("Balance updated! New amount: ", accountBalance)
	} else if choice == 3 {
		fmt.Print("Withdraw Amount: ")
		var withdrawAmount float64
		fmt.Scan(&withdrawAmount)
		
		if withdrawAmount <= 0 {
			fmt.Println("Invalid amount. Must be greater than zero")
			return
		}

		if withdrawAmount > accountBalance {
			fmt.Println("Invalid amount. You can't withdraw more than you have")
			return
		}

		accountBalance -= withdrawAmount
		fmt.Println("New balance: ", accountBalance)
	} else {
		fmt.Println("Goodbye!")
	}
}