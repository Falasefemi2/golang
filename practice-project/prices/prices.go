package prices

import (
	"fmt"

	"example.com/project/conversion"
	"example.com/project/iomanager"
)

type TaxIncludedPriceJob struct {
	IOManger          iomanager.IOManager `json:"-"`
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
}

func (job *TaxIncludedPriceJob) LoadData() error {
	// load file
	lines, err := job.IOManger.ReadLines()

	if err != nil {
		return err
	}

	prices, err := conversion.StringToFloat(lines)

	// Check if there's an error converting string to float64
	if err != nil {
		return err
	}
	// Store the price in the prices slice
	job.InputPrices = prices
	return nil
}

func (job *TaxIncludedPriceJob) Process() error {
	// Load data from the file
	err := job.LoadData()

	if err != nil {
		return err
	}

	// Create a map to store the results
	result := make(map[string]string)

	// Calculate tax-included prices for each price in InputPrices
	for _, price := range job.InputPrices {
		TaxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", TaxIncludedPrice)
	}

	// Print the results
	job.TaxIncludedPrices = result
	return job.IOManger.WriteResult(job)
}

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManger:    iom,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
