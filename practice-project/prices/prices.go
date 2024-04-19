package prices

import (
	"bufio"
	"fmt"
	"os"

	"example.com/project/conversion"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludedPriceJob) LoadData() {
	// Open the file named "prices.txt"
	file, err := os.Open("prices.txt")

	// Check if there's an error opening the file
	if err != nil {
		fmt.Println("An error occurred")
		fmt.Println(err)
		return
	}

	// Create a scanner to read from the file
	scanner := bufio.NewScanner(file)

	var lines []string

	// Read each line from the file
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Check if there's an error during scanning
	err = scanner.Err()

	if err != nil {
		fmt.Println("Reading the file failed")
		fmt.Println(err)
		file.Close()
		return
	}

	prices, err := conversion.StringToFloat(lines)

	// Check if there's an error converting string to float64
	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}
	// Store the price in the prices slice
	job.InputPrices = prices
	file.Close()

	// Store the prices in the InputPrices field of the struct
	job.InputPrices = prices

}

func (job *TaxIncludedPriceJob) Process() {

	// Load data from the file
	job.LoadData()

	// Create a map to store the results
	result := make(map[string]string)

	// Calculate tax-included prices for each price in InputPrices
	for _, price := range job.InputPrices {
		TaxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", TaxIncludedPrice)
	}

	// Print the results
	fmt.Println(result)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
