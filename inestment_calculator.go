package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const (
	accountBalanceFile  = "balance.txt"
	defaultBalanceValue = 1000
)

func writeFloatToFile(value float64, fileName string) {
	valueTxt := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueTxt), 0o644)
}

func getFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 1000, errors.New("failed to find file")
	}
	valueTxt := string(data)
	valueNum, err := strconv.ParseFloat(valueTxt, 64)
	if err != nil {
		return 1000, errors.New("failed to parse stored value")
	}
	return valueNum, nil
}

func main() {
	accountBalance, err := getFloatFromFile(accountBalanceFile)
	if err != nil {
		fmt.Println("ERROR: ")
		fmt.Println(err)
		panic("Can't continue, sorry")
	}

	var choice int

	presentOptions()

	for choice != 4 {
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)
		fmt.Println("Your choice is: ", choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is: ", accountBalance)
		case 2:
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
			writeFloatToFile(accountBalance, accountBalanceFile)
		case 3:
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
			writeFloatToFile(accountBalance, accountBalanceFile)

		case 4:
			fmt.Println("Goodbye!")
			return

		default:

		}
	}
	fmt.Println("Thanks for choosing our Bank!")
}
