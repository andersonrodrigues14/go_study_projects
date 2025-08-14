package main

import (
	"fmt"
)

func main() {
	prices := []float64{10.99, 8.99}

	fmt.Println(prices[1])
	prices[1] = 9.9

	prices = append(prices, 5.99)

	fmt.Println(prices)

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
