package main

import "fmt"

func main() {
	fmt.Println("Welcome to Go Brank!")
	fmt.Println("What do tou want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Println("Your choice: ")
	fmt.Scan(&choice)
	fmt.Println("Your choice: ", choice)
}
