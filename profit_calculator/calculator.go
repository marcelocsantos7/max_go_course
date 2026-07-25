package main

import (
	"errors"
	"fmt"
	"os"
)

func Results(revenue, expenses, taxrate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxrate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("Value must be positive")
	}
	return userInput, nil
}

func storeResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %1.f\nRATIO: %.1f", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
}

func main() {
	revenue, err := getUserInput("Revenue: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	expenses, err := getUserInput("Expenses: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	taxRate, err := getUserInput("Tax Rate: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	ebt, profit, ratio := Results(revenue, expenses, taxRate)

	fmt.Printf("Earnings: %.1f\n", ebt)
	fmt.Printf("Profit: %.1f\n", profit)
	fmt.Printf("Ratio: %.1f\n", ratio)
	storeResults(ebt, profit, ratio)
}
