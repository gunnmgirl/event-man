package main

import "fmt"

func main() {
	userNames := make([]string, 2, 5)

	userNames[0] = "Michael"
	userNames = append(userNames, "Dzana")
	userNames = append(userNames, "Nikita")

	fmt.Println("userNames", userNames)

	courseRatings := make(map[string]float64, 2)
	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8

	fmt.Println("courseRatings", courseRatings)
}
