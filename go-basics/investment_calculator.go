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

	// add string in variabel
	formattedFV := fmt.Sprintf("Futere Value: %.2f\n", futereValeu)
	formattedRFV := fmt.Sprintf("Futere Value (adjusted for Inflation): %.2f",futereRealValue)

	//output information

	//fmt.Println("Futere Value: ",futereValeu)
	//fmt.Println("Futere Value (adjusted for Inflation): ",futereRealValue)
	//fmt.Printf("Futere Value: %.2f\n", futereValeu)
	//fmt.Printf("Futere Value (adjusted for Inflation): %.2f",futereRealValue)
	fmt.Print(formattedFV, formattedRFV)
}