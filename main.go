/**
 * author:
 * version: 1.0.0
 * date: 2025-12-09
 * fileoverview: A program that calculates how many years it will take for a bank account
 */

package main

import ("fmt")

func main() {

	// variables
	var startingAmount float64 = 0
	var interestRate float64 = 0
	var neededAmount float64 = 0
	var years int = 0

	// input from the user
	fmt.Print("Enter the starting bank account amount: ")
	fmt.Scan(&startingAmount)

	fmt.Print("Enter the yearly interest rate (as a percentage): ")
	fmt.Scan(&interestRate)

	fmt.Print("Enter the amount needed for post-secondary education: ")
	fmt.Scan(&neededAmount)

	// calculate the number of years using a while-like loop
	currentAmount := startingAmount

	for currentAmount < neededAmount {
		currentAmount = currentAmount + (currentAmount * (interestRate / 100))
		years = years + 1
	}

	// output
	fmt.Printf("It will take %d years for the starting bank account to reach the required amount for post-secondary education.\n", years)
	fmt.Println("Done.")
}