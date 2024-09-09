package main

import (
	"fmt"
	"math"
)

// Format as currency
func formatCurrency(amount float64) string {
	return fmt.Sprintf("$%.2f", amount)
}

func calculateFutureValue(investmentAmount, expectedReturnRate float64, years int, annualContribution float64) float64 {
	futureValue := investmentAmount
	for i := 1; i <= years; i++ {
		// Calculate interest
		futureValue += futureValue * (expectedReturnRate / 100)
		// Add annual contributions
		futureValue += annualContribution
		fmt.Printf("Year %d: %s\n", i, formatCurrency(futureValue))
	}
	return futureValue
}

func main() {
	const inflationRate = 2.5
	var investmentAmount, expectedReturnRate, annualContribution float64
	var years int

	fmt.Println("Answer the questions and get your future value calculated with yearly breakdown!")

	// Prompt the user for input
	fmt.Print("Enter the initial investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter the expected annual return rate (in %): ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Enter the number of years: ")
	fmt.Scan(&years)

	fmt.Print("Enter the annual contribution amount: ")
	fmt.Scan(&annualContribution)

	// Calculate future value
	futureValue := calculateFutureValue(investmentAmount, expectedReturnRate, years, annualContribution)

	futureRealValue := futureValue / math.Pow(1+inflationRate/100, float64(years))

	// Display final value
	fmt.Println("\nFinal Future Value:")
	fmt.Println(formatCurrency(futureValue))
	fmt.Println("Future real value: ", futureRealValue)
}
