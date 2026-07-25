package main

import (
	"fmt"
)

func main () {
	revenue := Revenue()
	expenses := Expenses()
	taxRate := TaxRate()

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)
}

func Revenue() float64 {
	var revenue float64
	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)
	return revenue
}

func Expenses() float64 {
	var expenses float64
	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)
	return expenses
}

func TaxRate() float64 {
	var taxRate float64
	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)
	return taxRate
}