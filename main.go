package main

import "fmt"

func main() {
	fmt.Println(sumup(5, 5, 5))
}

func sumup(numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum = sum + val
	}
	return sum
}
