package prices

import (
	"fmt"

	"example.com/project/conversion"
	"example.com/project/filemanager"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludedPriceJob) LoadData() {
	// load file
	lines, err := filemanager.ReadLines("prices.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	prices, err := conversion.StringToFloat(lines)

	// Check if there's an error converting string to float64
	if err != nil {
		fmt.Println(err)
		return
	}
	// Store the price in the prices slice
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
