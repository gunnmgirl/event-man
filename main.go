package main

import "fmt"

func main() {
	fmt.Println(sumup(1, 5, 5, 5))
}

func sumup(startingNum int, numbers ...int) int {
	sum := startingNum

	for _, val := range numbers {
		sum = sum + val
	}
	return sum
}
