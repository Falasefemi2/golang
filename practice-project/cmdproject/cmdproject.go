package cmdproject

import "fmt"

type CMDManger struct {
}

func (cmd CMDManger) ReadLines() ([]string, error) {
	fmt.Println("Please enter your prices. Confirm every price with ENTER")

	var prices []string

	for {
		var price string
		fmt.Print("Prices: ")
		fmt.Scan(&price)

		if price == "0" {
			break
		}
		prices = append(prices, price)
	}

	return prices, nil
}

func (cmd CMDManger) WriteResult(data any) error {
	fmt.Println(data)
	return nil
}

func New() CMDManger {
	return CMDManger{}
}
