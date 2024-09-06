package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64
	var expectedReturnRate float64
	var years int

	fmt.Println("Answer the questions and get your future value calculated")
	// Prompt the user for input
	fmt.Print("Enter the investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter the expected return rate (in %): ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Enter the number of years: ")
	fmt.Scan(&years)

	futureValue := float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))
	fmt.Println("Amazing, your future value has been calculated")
	fmt.Println(futureValue)
}
