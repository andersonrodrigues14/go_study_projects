package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, err := getUserInput("Revenue: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Print("Revenue: ")
	// fmt.Scan(&revenue)

	expenses, err := getUserInput("Expenses: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Print("Expenses: ")
	// fmt.Scan(&expenses)

	taxRate, err := getUserInput("Taxt Rate: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	// if err1 != nil || err2 != nil || err3 != nil {
	// 	fmt.Println(err1)
	// 	return
	// }
	// fmt.Print("Taxt Rate: ")
	// fmt.Scan(&taxRate)

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)
	// ebt := revenue - expenses
	// profit := ebt * (1 - taxRate/100)
	// ratio := ebt / profit

	fmt.Printf("%.2f\n", ebt)
	fmt.Printf("%.2f\n", profit)
	fmt.Printf("%.2f\n", ratio)
	storeResults(ebt, profit, ratio)
}

func storeResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.2f\nProfif: %.2f\nRatio: %.2f\n", ebt, profit, ratio)
	os.WriteFile("results.tx", []byte(results), 0644)
}

func calculateFinancials(revenue float64, expenses float64, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	return ebt, profit, ratio
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("value must be a positive number")
	}

	return userInput, nil
}
