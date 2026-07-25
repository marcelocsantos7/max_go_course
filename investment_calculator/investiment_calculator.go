package main

import (
	"fmt"
	"math"
)

func main() {
	calc()
}

func calc() {
	var amount float64
	var expect_rate float64
	var years float64
	// amount, expect_rate, years := 1000.0, 12.5, 10.0

	const inflationRate = 4.5

	fmt.Print("Please, inform the amount to invest: ")
	fmt.Scan(&amount)
	fmt.Print("Please, inform the expect rate: ")
	fmt.Scan(&expect_rate)
	fmt.Print("How long? (years): ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(amount, expect_rate, years, inflationRate)

	formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedFRV := fmt.Sprintf("Future Real Value: %.2f\n", futureRealValue)

	fmt.Println("Investment amount: ", amount)
	fmt.Println("Expect return rate: ", expect_rate)
	fmt.Println("Years: ", years)

	fmt.Printf("%v \n %v", formattedFV, formattedFRV)

}

func calculateFutureValues(amount, rate, years, inflation float64) (fv float64, rfv float64) {
	fv = amount * math.Pow(1+rate/100, years)
	rfv = fv / math.Pow(1+inflation/100, years)
	return fv, rfv
	// return
}
