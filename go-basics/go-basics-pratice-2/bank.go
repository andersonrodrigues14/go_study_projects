package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return 1000, errors.New("failed to find balance file")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("failed to parse stored balance value")
	}

	return balance, nil
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {

	for {
		var accountBalance, err = getBalanceFromFile()

		if err != nil {
			fmt.Println("ERROR")
			fmt.Println(err)
			fmt.Println("----------")
			//return
			//panic("Can't continue, sorry.") // finished aplication
		}

		fmt.Println("Welcome to Go Brank!")

		fmt.Println("What do tou want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is: ", accountBalance)
		case 2:
			fmt.Print("Your deposit:")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0")
				//return
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Balance updated! New amount: ", accountBalance)
			writeBalanceToFile(accountBalance)
		case 3:
			fmt.Print("Withdrawal amount:")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0")
				//return
				continue
			}

			if withdrawalAmount > accountBalance {
				fmt.Println("Invalid amount. Youn can't withdraw more than you have.")
				//return
				continue
			}

			accountBalance -= withdrawalAmount
			fmt.Println("Balance updated! New amount: ", accountBalance)
			writeBalanceToFile(accountBalance)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosin our bank")
			return
			//break
		}
	}

	// for {

	// 	fmt.Println("What do tou want to do?")
	// 	fmt.Println("1. Check balance")
	// 	fmt.Println("2. Deposit money")
	// 	fmt.Println("3. Withdraw money")
	// 	fmt.Println("4. Exit")

	// 	var choice int
	// 	fmt.Print("Your choice: ")
	// 	fmt.Scan(&choice)

	// 	if choice == 1 {
	// 		fmt.Println("Your balance is: ", accountBalance)
	// 	} else if choice == 2 {
	// 		fmt.Print("Your deposit:")
	// 		var depositAmount float64
	// 		fmt.Scan(&depositAmount)

	// 		if depositAmount <= 0 {
	// 			fmt.Println("Invalid amount. Must be greater than 0")
	// 			//return
	// 			continue
	// 		}

	// 		accountBalance += depositAmount
	// 		fmt.Println("Balance updated! New amount: ", accountBalance)
	// 	} else if choice == 3 {
	// 		fmt.Print("Withdrawal amount:")
	// 		var withdrawalAmount float64
	// 		fmt.Scan(&withdrawalAmount)

	// 		if withdrawalAmount <= 0 {
	// 			fmt.Println("Invalid amount. Must be greater than 0")
	// 			//return
	// 			continue
	// 		}

	// 		if withdrawalAmount > accountBalance {
	// 			fmt.Println("Invalid amount. Youn can't withdraw more than you have.")
	// 			// return
	// 			continue
	// 		}

	// 		accountBalance -= withdrawalAmount
	// 		fmt.Println("Balance updated! New amount: ", accountBalance)
	// 	} else {
	// 		fmt.Println("Goodbye!")
	// 		//return
	// 		break
	// 	}

	// 	fmt.Println("Thanks for choosin our bank")
	// }

}
