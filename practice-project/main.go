package main

import (
	"example.com/practice-project/prices"
)

func main() {
	taxes := []float64{0, 0.07, 0.1, 0.15}

	for _, tax := range taxes {
		priceJob := prices.NewTaxIncludedPriceJob(tax)
		priceJob.Process()
	}
}
