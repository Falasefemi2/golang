package main

import "fmt"

func main() {
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[0:1])
	prices[1] = 9.99
	updatedPrices := append(prices, 5.99)
	fmt.Println(updatedPrices, prices)
}

// func main() {
// 	var productNames [4]string = [4]string{"A book"}
// 	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
// 	fmt.Println(prices)

// 	productNames[2] = "A Carpet"

// 	fmt.Println(productNames)
// 	fmt.Println(prices[2])
// 	fmt.Println(productNames)

// 	featuredPrices := prices[1:]
// 	featuredPrices[0] = 199.99
// 	highLightedPrices := featuredPrices[:1]
// 	fmt.Println(highLightedPrices)
// 	fmt.Println(featuredPrices)
// 	fmt.Println(len(highLightedPrices), cap(highLightedPrices))

// 	highLightedPrices = highLightedPrices[:3]
// 	fmt.Println(highLightedPrices)
// 	fmt.Println(featuredPrices)
// 	fmt.Println(len(highLightedPrices), cap(highLightedPrices))
// }
