package main

import "fmt"

func main() {
	var accountBalance float64 = 1000

	fmt.Println("Welcome to Go Brank!")

	for {

		fmt.Println("What do tou want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("Your balance is: ", accountBalance)
		} else if choice == 2 {
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
		} else if choice == 3 {
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
				// return
				continue
			}

			accountBalance -= withdrawalAmount
			fmt.Println("Balance updated! New amount: ", accountBalance)
		} else {
			fmt.Println("Goodbye!")
			//return
			break
		}

		fmt.Println("Thanks for choosin our bank")
	}

}
