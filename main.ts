/**
 * @author 
 * @version 1.0.0
 * @date 2025-12-09
 * @fileoverview A program that calculates how many years it will take for a bank account
 */

// variables
let startingAmount: number = 0;
let interestRate: number = 0;
let neededAmount: number = 0;
let years: number = 0;

// input from the user
startingAmount = Number(prompt("Enter the starting bank account amount: ") || "0");
interestRate = Number(prompt("Enter the yearly interest rate (as a percentage): ") || "0");
neededAmount = Number(prompt("Enter the amount needed for post-secondary education: ") || "0");

// calculate the number of years using a while loop
let currentAmount: number = startingAmount;

while (currentAmount < neededAmount) {
  currentAmount = currentAmount + (currentAmount * (interestRate / 100));
  years = years + 1;
}

// output
console.log("It will take " + years + " years for the starting bank account to reach the required amount for post-secondary education.");
console.log("Done.");