package prices

import (
	"fmt"

	"example.com/project/conversion"
	"example.com/project/filemanager"
)

type TaxIncludedPriceJob struct {
	IOManger          filemanager.FileManager `json:"-"`
	TaxRate           float64                 `json:"tax_rate"`
	InputPrices       []float64               `json:"input_prices"`
	TaxIncludedPrices map[string]string       `json:"tax_included_prices"`
}

func (job *TaxIncludedPriceJob) LoadData() {
	// load file
	lines, err := job.IOManger.ReadLines()

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
	job.TaxIncludedPrices = result

	job.IOManger.WriteResult(job)
}

func NewTaxIncludedPriceJob(fm filemanager.FileManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManger:    fm,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
