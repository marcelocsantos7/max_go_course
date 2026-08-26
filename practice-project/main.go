package main

import (
	"fmt"
)

func main() {
	prices := []float64{10, 20, 30}
	taxes := []float64{0, 0.07, 0.1, 0.15}

	result := make(map[float64][]float64)

	for _, tax := range taxes {
		taxedPrices := make([]float64, len(prices))
		for idx, price := range prices {
			taxedPrices[idx] = price * (1 + tax)
		}
		result[tax] = taxedPrices
	}
	fmt.Println(result)
}
