package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const (
	accountBalanceFile  = "balance.txt"
	defaultBalanceValue = 1000
)

func writeBalanceToFile(balance float64) {
	balanceTxt := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceTxt), 0o644)
}

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		return defaultBalanceValue, errors.New("Failed to find balance file")
	}
	balanceTxt := string(data)
	balanceNum, err := strconv.ParseFloat(balanceTxt, 64)
	if err != nil {
		return defaultBalanceValue, errors.New("Failed to parse stored balance value")
	}
	return balanceNum, nil
}

// const taxRate = 19 // in percentage

// // Format as currency
// func formatCurrency(amount float64) string {
// 	return fmt.Sprintf("$%.2f", amount)
// }

// func calculateEarnings(revenueVal, ExpensesVal, taxRateVal float64) (float64, float64, float64) {
// 	earningsBeforeTax := revenueVal - ExpensesVal
// 	taxAmount := earningsBeforeTax * (taxRateVal / 100)
// 	earningsAfterTax := earningsBeforeTax - taxAmount
// 	ratio := earningsBeforeTax / earningsAfterTax

// 	return earningsBeforeTax, earningsAfterTax, ratio
// }

// func enterUserValue(text string, variable *float64) {
// 	fmt.Printf("Enter the %s: ", text)
// 	fmt.Scan(variable)
// }

func main() {
	accountBalance, err := getBalanceFromFile()
	if err != nil {
		fmt.Println("ERROR: ")
		fmt.Println(err)
		panic("Can't continue, sorry")
	}

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	var choice int

	for choice != 4 {
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)
		fmt.Println("Your choice is: ", choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is: ", accountBalance)
		case 2:
			fmt.Print("Enter an amount to deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			for depositAmount <= 0 {
				fmt.Println("You need to enter a value greater then zero to deposit money")
				fmt.Print("Enter new deposit amount: ")
				fmt.Scan(&depositAmount)
			}
			accountBalance += depositAmount
			fmt.Println("New account balance is: ", accountBalance)
			writeBalanceToFile(accountBalance)
		case 3:
			fmt.Print("Enter an amount to withdraw: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)
			for withdrawAmount > accountBalance || withdrawAmount <= 0 {
				fmt.Println("You surpassed your account balance! Please withdraw lesser amount of money")
				fmt.Print("Enter new withdraw amount: ")
				fmt.Scan(&withdrawAmount)
			}
			accountBalance -= withdrawAmount
			fmt.Println("New account balance is: ", accountBalance)
			writeBalanceToFile(accountBalance)

		case 4:
			fmt.Println("Goodbye!")
			return

		}
	}
	fmt.Println("Thanks for choosing our Bank!")
}
