package main

import (
	"fmt"

	"example.com/project/filemanager"
	"example.com/project/prices"
)

func main() {
	taxRate := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRate))

	for index, taxRate := range taxRate {
		doneChans[index] = make(chan bool)
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// cmdm := cmdproject.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.Process(doneChans[index])

		// if err != nil {
		// 	fmt.Println("could not process job")
		// 	fmt.Println(err)
		// }
	}

}
