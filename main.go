package main

import "fmt"

func main() {
	userNames := make([]string, 2, 5)

	userNames[0] = "Michael"
	userNames = append(userNames, "Dzana")
	userNames = append(userNames, "Nikita")

	fmt.Println("userNames", userNames)
}