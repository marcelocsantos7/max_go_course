package main

import (
	"bank/fileops"
	"fmt"
)

const accountBalanceFile = "balance.txt"

func Balance(ac_balance float64) {
	fmt.Printf("\nYour balance is %.2f\n", ac_balance)
}

func Withdraw(value float64, balance float64) float64 {
	if value <= balance {
		balance -= value
		fmt.Printf("Withdraw value: %.2f successfull\n", value)
		fmt.Printf("Balance: %.2f\n", balance)
		return balance
	}
	fmt.Println("You don't have enough funds")
	return balance
}

func Deposit(value float64, accountBallance float64) float64 {
	newBallance := accountBallance + value
	fmt.Printf("Balance: %.2f\n", newBallance)
	return newBallance
}

func main() {
	accountBallance, err := fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("------------")
		fmt.Println(err)
		fmt.Println("------------")
		// panic("Can't continue. Sorry.")
	}

	fmt.Println("Welcome to Go Bank!")
	for {
		presentOptions()
		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			Balance(accountBallance)

		case 2:
			var deposit float64
			fmt.Print("Deposit value: ")
			fmt.Scan(&deposit)
			if deposit <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}
			accountBallance = Deposit(deposit, accountBallance)
			if err := fileops.WriteFloatToFile(accountBallance, accountBalanceFile); err != nil {
				fmt.Println(err)
			}

		case 3:
			fmt.Print("Witdraw value: ")
			var withdraw float64
			fmt.Scan(&withdraw)
			accountBallance = Withdraw(withdraw, accountBallance)
			if err := fileops.WriteFloatToFile(accountBallance, accountBalanceFile); err != nil {
				fmt.Println(err)
			}

		default:
			fmt.Println("Session End")
			return
		}

	}
}
