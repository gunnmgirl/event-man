package main

import (
	"fmt"
)

// Format as currency
func formatCurrency(amount float64) string {
	return fmt.Sprintf("$%.2f", amount)
}

func calculateEarnings(revenueVal, ExpensesVal, taxRateVal float64) (float64, float64, float64) {
	earningsBeforeTax := revenueVal - ExpensesVal
	taxAmount := earningsBeforeTax * (taxRateVal / 100)
	earningsAfterTax := earningsBeforeTax - taxAmount
	ratio := earningsBeforeTax / earningsAfterTax

	return earningsBeforeTax, earningsAfterTax, ratio
}

func main() {
	const taxRate = 19 // in percentage
	var revenue, expenses float64

	fmt.Println("Answer the questions and get value of your earnings before and after taxes")

	fmt.Print("Enter the revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter the expected expenses: ")
	fmt.Scan(&expenses)

	// Calculate earnings
	ebt, profit, ratio := calculateEarnings(revenue, expenses, taxRate)

	fmt.Println("Earnings before tax are: ", formatCurrency(ebt))
	fmt.Println("Earnings after tax are: ", formatCurrency(profit))
	fmt.Printf("Ratio EBT/profit is: %.2f\n", ratio)
}
