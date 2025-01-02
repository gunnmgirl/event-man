package main

import "fmt"

func main() {
	fmt.Println(factorial((5)))
}

func factorial(number int) int {
	result := 1
	for i := 1; i <= number; i++ {
		result = result * i
	}
	return result
}
