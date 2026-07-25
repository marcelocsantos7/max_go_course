package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
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

func writeBalanceToFile(balance float64) error {
	balanceText := fmt.Sprint(balance)

	// Write the balance to a file in the current working directory
	filePath := filepath.Join(".", "balance.txt")

	// Write file (0644 gives read/write permissions to owner, read to others)
	err := os.WriteFile(filePath, []byte(balanceText), 0644)
	if err != nil {
		return fmt.Errorf("failed to write balance to file: %w", err)
	}

	return nil
}

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile("balance.txt")
	if err != nil {
		return 0, errors.New("Failed to read balance file")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 0, errors.New("Failed to parse file value")
	}
	return balance, nil
}

func main() {
	accountBallance, err := getBalanceFromFile()

	if err != nil {
		fmt.Println("------------")
		fmt.Println(err)
		fmt.Println("------------")
		// panic("Can't continue. Sorry.")
	}

	fmt.Println("Welcome to Go Bank!")
	for {
		fmt.Println("\nWhat do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Printf("4. Exit\n\n")

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
			if err := writeBalanceToFile(accountBallance); err != nil {
				fmt.Println(err)
			}

		case 3:
			fmt.Print("Witdraw value: ")
			var withdraw float64
			fmt.Scan(&withdraw)
			accountBallance = Withdraw(withdraw, accountBallance)
			if err := writeBalanceToFile(accountBallance); err != nil {
				fmt.Println(err)
			}

		default:
			fmt.Println("Session End")
			return
		}

	}
}
