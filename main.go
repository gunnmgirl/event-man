package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4}

	doubles := transformNumbers(&numbers, double)
	fmt.Println("doubles-", doubles)
	triples := transformNumbers(&numbers, triple)
	fmt.Println("triples-", triples)
}

func transformNumbers(numbers *[]int, transformFunc func(int) int) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transformFunc(val))
	}
	return dNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
