package main

const (
	accountBalanceFile  = "balance.txt"
	defaultBalanceValue = 1000
)

// func main() {
// 	accountBalance, err := fileops.GetFloatFromFile(accountBalanceFile)
// 	if err != nil {
// 		fmt.Println("ERROR: ")
// 		fmt.Println(err)
// 		panic("Can't continue, sorry")
// 	}

// 	var choice int

// 	presentOptions()

// 	for choice != 4 {
// 		fmt.Print("Enter your choice: ")
// 		fmt.Scan(&choice)
// 		fmt.Println("Your choice is: ", choice)

// 		switch choice {
// 		case 1:
// 			fmt.Println("Your balance is: ", accountBalance)
// 		case 2:
// 			fmt.Print("Enter an amount to deposit: ")
// 			var depositAmount float64
// 			fmt.Scan(&depositAmount)
// 			for depositAmount <= 0 {
// 				fmt.Println("You need to enter a value greater then zero to deposit money")
// 				fmt.Print("Enter new deposit amount: ")
// 				fmt.Scan(&depositAmount)
// 			}
// 			accountBalance += depositAmount
// 			fmt.Println("New account balance is: ", accountBalance)
// 			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
// 		case 3:
// 			fmt.Print("Enter an amount to withdraw: ")
// 			var withdrawAmount float64
// 			fmt.Scan(&withdrawAmount)
// 			for withdrawAmount > accountBalance || withdrawAmount <= 0 {
// 				fmt.Println("You surpassed your account balance! Please withdraw lesser amount of money")
// 				fmt.Print("Enter new withdraw amount: ")
// 				fmt.Scan(&withdrawAmount)
// 			}
// 			accountBalance -= withdrawAmount
// 			fmt.Println("New account balance is: ", accountBalance)
// 			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)

// 		case 4:
// 			fmt.Println("Goodbye!")
// 			return

// 		default:

// 		}
// 	}
// 	fmt.Println("Thanks for choosing our Bank!")
// }
