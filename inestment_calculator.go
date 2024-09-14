package main

import (
	"fmt"
)

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
	// var revenue, expenses float64

	// fmt.Println("Answer the questions and get value of your earnings before and after taxes")

	// enterUserValue("revenue", &revenue)
	// enterUserValue("expected expenses", &expenses)

	// // Calculate earnings
	// ebt, profit, ratio := calculateEarnings(revenue, expenses, taxRate)

	// fmt.Println("Earnings before tax are: ", formatCurrency(ebt))
	// fmt.Println("Earnings after tax are: ", formatCurrency(profit))
	// fmt.Printf("Ratio EBT/profit is: %.2f\n", ratio)
	accountBalance := 1000.00

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
		if choice == 1 {
			fmt.Println("Your balance is: ", accountBalance)
			continue
		} else if choice == 2 {
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
			continue
		} else if choice == 3 {
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
			continue
		} else if choice == 4 {
			fmt.Println("Goodbye!")
			break
		}
	}
	fmt.Println("Thanks for choosing our Bank!")
}
