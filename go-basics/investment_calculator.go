package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Investment Years: ")
	fmt.Scan(&years)

	futereValeu := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futereRealValue := futereValeu / math.Pow(1+inflationRate/100, years)

	fmt.Println(futereValeu)
	fmt.Println(futereRealValue)
}
