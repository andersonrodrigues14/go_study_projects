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

	//Closures
	double := createTransformer(2)
	doubled2 := transformNumbers(&numbers, double)
	fmt.Println("Closures: ", doubled2)

	// Recursion Functions
	fact := factorial(5)
	fmt.Println(fact)

}

// func transformNumbers(numbers *[]int, transform transformFn) []int {
// 	dNumbers := []int{}
// 	for _, val := range *numbers {
// 		dNumbers = append(dNumbers, transform(val))
// 	}
// 	return dNumbers
// }

func factorial(number int) int {
	if number == 0 {
		return 1
	}
	return number * factorial(number - 1)

	// result := 1

	// for i:= 1; i <= number; i++ {
	// 	result = result * i
	// }

	// return result
}

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

func createTransformer(factor int) func(int) int{
	return func(number int) int{
		return number * factor
	}
}