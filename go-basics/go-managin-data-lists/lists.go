package main

import (
	"fmt"
)

func main() {
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[0:1])
	prices[1] = 9.9

	prices = append(prices, 5.99, 12.99, 29.99, 100.10)
	prices = prices[1:]
	fmt.Println(prices)

	discountPrices := []float64{101.99, 80.99, 20.59}
	prices = append(prices, discountPrices...)
	fmt.Println(prices)

	userNames := make([]string, 2, 5)
	userNames[0] = "Julie"
	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel")
	fmt.Println(userNames)

	for index, value := range userNames {
		fmt.Println(index)
		fmt.Println(value)
	}
}

// func main() {
// 	var productNames [4]string = [4]string{"A book"}
// 	productNames[2] = "A Carpet"

// 	prices := [4]float64{10.96, 9.99, 45.99, 20.0}

// 	fmt.Println(prices)
// 	fmt.Println(productNames)

// 	featurePrices := prices[1:3] // slices
// 	fmt.Println(featurePrices)
// 	fmt.Println(len(featurePrices)) // lenth of array
// 	fmt.Println(cap(featurePrices)) // capacity of array
// }
