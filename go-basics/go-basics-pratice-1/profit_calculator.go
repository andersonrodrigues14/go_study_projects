package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	revenue = getUserInput("Revenue: ")
	// fmt.Print("Revenue: ")
	// fmt.Scan(&revenue)

	expenses = getUserInput("Expenses: ")
	// fmt.Print("Expenses: ")
	// fmt.Scan(&expenses)

	taxRate = getUserInput("Taxt Rate: ")
	// fmt.Print("Taxt Rate: ")
	// fmt.Scan(&taxRate)

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)
	// ebt := revenue - expenses
	// profit := ebt * (1 - taxRate/100)
	// ratio := ebt / profit

	fmt.Printf("%.2f\n", ebt)
	fmt.Printf("%.2f\n", profit)
	fmt.Printf("%.2f\n", ratio)
}

func calculateFinancials(revenue float64, expenses float64, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	return ebt, profit, ratio
}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}
