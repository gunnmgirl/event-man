package main

import "fmt"

func main() {
	prices := []float64{10.09, 11.33} // slice
	fmt.Println("first-", prices[0])
	newSlice := append(prices, 5.88)
	fmt.Println("newSlice-", newSlice)
}

// func main() {
// 	prices := [4]float64{10.99, 9.99, 45.99, 20.00}
// 	fmt.Println("prices-", prices[3])

// 	featuredPrices := prices[:1] // second value is excluded
// 	fmt.Println("featuredPrices-", featuredPrices)

// 	fmt.Println("len and cap -", len(featuredPrices), cap(featuredPrices))
// }
