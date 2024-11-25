package main

import "fmt"

func main() {
	prices := [4]float64{10.99, 9.99, 45.99, 20.00}
	fmt.Println("prices-", prices[3])

	featuredPrices := prices[1:3] // second value is excluded
	fmt.Println("featuredPrices-", featuredPrices)
}
