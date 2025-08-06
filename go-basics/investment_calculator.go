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

	//fmt.Print("Investment Amount: ")
	outputTex("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Investment Years: ")
	fmt.Scan(&years)

	futereValeu, futereRealValue := calculateFutereValues(investmentAmount, expectedReturnRate, years, inflationRate)

	//futereValeu := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	//futereRealValue := futereValeu / math.Pow(1+inflationRate/100, years)

	// add string in variabel
	formattedFV := fmt.Sprintf("Futere Value: %.2f\n", futereValeu)
	formattedRFV := fmt.Sprintf("Futere Value (adjusted for Inflation): %.2f", futereRealValue)

	//output information

	//fmt.Println("Futere Value: ",futereValeu)
	//fmt.Println("Futere Value (adjusted for Inflation): ",futereRealValue)
	//fmt.Printf("Futere Value: %.2f\n", futereValeu)
	//fmt.Printf("Futere Value (adjusted for Inflation): %.2f",futereRealValue)

	//More lines
	// fmt.Printf(`
	// Futere Value: %.2f
	// Futere Value (adjusted for Inflation): %.2f
	// ` ,futereRealValue, futereRealValue)

	fmt.Print(formattedFV, formattedRFV)
}

func outputTex(text string) {
	fmt.Print(text)
}

func calculateFutereValues(investmentAmount float64, expectedReturnRate float64, years float64, inflationRate float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv := fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}
