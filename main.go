package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println("map-", m)
}

func main() {
	userNames := make([]string, 2, 5)

	userNames[0] = "Michael"
	userNames = append(userNames, "Dzana")
	userNames = append(userNames, "Nikita")

	fmt.Println("userNames", userNames)

	courseRatings := make(floatMap, 2)
	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8

	courseRatings.output()
}
