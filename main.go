package main

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}

	doubles := transformNumbers(&numbers, getTransformerFunction(&numbers))
	fmt.Println("doubles-", doubles)
	triples := transformNumbers(&numbers, getTransformerFunction(&numbers))
	fmt.Println("triples-", triples)
}

func transformNumbers(numbers *[]int, transformFunc transformFn) []int {
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

func getTransformerFunction(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	}
	return triple
}
