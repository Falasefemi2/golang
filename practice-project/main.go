package main

import (
	"example.com/project/cmdproject"
	"example.com/project/prices"
)

func main() {
	taxRate := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRate {
		// fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		cmdm := cmdproject.New()
		priceJob := prices.NewTaxIncludedPriceJob(cmdm, taxRate)
		priceJob.Process()
	}

}
