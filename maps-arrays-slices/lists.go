package main

import (
	"fmt"
)

func main() {
	fmt.Println("Slices, Maps and Arrays")

	var productNames [4]string = [4]string{"A Book"}

	prices := [4]float64{11.1, 22.2, 33.33, 44.44}
	fmt.Println("prices original ", prices)

	productNames[2] = "A Carpet"

	fmt.Println(prices[0])
	fmt.Println(productNames)

	pricesSlice := prices[1:3] // slice: first element included, last excluded
	featuredPrices := prices[1:]
	featuredPrices[0] = 199.99
	highlightedPrices := featuredPrices[1:]

	fmt.Println("prices slice prices[1:3]:", pricesSlice)

	fmt.Println("Featured ", featuredPrices)
	fmt.Println("prices updated", prices)
	fmt.Println("Highlighted", highlightedPrices)
	fmt.Println(len(featuredPrices), cap(featuredPrices))
}
