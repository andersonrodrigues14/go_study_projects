package main

import (
	"fmt"
)

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)

	moreNumbers := []int{5,1,2}
	transformerFn1 := getTransformerFunction(&numbers)
	transformerFn2 := getTransformerFunction(&moreNumbers)

	transformedNumbers := transformNumbers(&numbers, transformerFn1)
	moreTransformedNumbers := transformNumbers(&moreNumbers, transformerFn2)
	fmt.Println(transformedNumbers)
	fmt.Println(moreTransformedNumbers)

	//Anonymous Functions
	transformed := transformNumbers(&numbers, func(number int)  int {
		return number * 2
	})

	fmt.Println("Anonymous Functions: ", transformed)

}

// func transformNumbers(numbers *[]int, transform transformFn) []int {
// 	dNumbers := []int{}
// 	for _, val := range *numbers {
// 		dNumbers = append(dNumbers, transform(val))
// 	}
// 	return dNumbers
// }

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}
	return dNumbers
}

func getTransformerFunction(numbers *[]int) transformFn{
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
