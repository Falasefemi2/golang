package main

import (
	"fmt"

	"example.com/project/filemanager"
	"example.com/project/prices"
)

func main() {
	taxRate := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRate {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// cmdm := cmdproject.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		err := priceJob.Process()

		if err != nil {
			fmt.Println("could not process job")
			fmt.Println(err)
		}
	}

}
