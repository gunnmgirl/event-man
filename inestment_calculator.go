package main

import (
	"fmt"
)

const taxRate = 19 // in percentage

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

func enterUserValue(text string, variable *float64) {
	fmt.Printf("Enter the %s: ", text)
	fmt.Scan(variable)
}

func main() {
	var revenue, expenses float64

	fmt.Println("Answer the questions and get value of your earnings before and after taxes")

	enterUserValue("revenue", &revenue)
	enterUserValue("expected expenses", &expenses)

	// Calculate earnings
	ebt, profit, ratio := calculateEarnings(revenue, expenses, taxRate)

	fmt.Println("Earnings before tax are: ", formatCurrency(ebt))
	fmt.Println("Earnings after tax are: ", formatCurrency(profit))
	fmt.Printf("Ratio EBT/profit is: %.2f\n", ratio)
}
