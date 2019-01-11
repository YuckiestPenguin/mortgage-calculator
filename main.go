package main

import (
	"fmt"
	"math"
)

const vaFundingFee float64 = 0.0215
const lengthOfLoan float64 = 360

var housePrice float64
var interestRate float64
var propertyTax float64
var hoa float64
var homeInsurance float64

func main() {
	fmt.Println("How much is the house?")
	fmt.Scan(&housePrice)

	fmt.Println("What is your current interest rate?")
	fmt.Scan(&interestRate)

	fmt.Println("How much will your ANNUAL Property Tax be?")
	fmt.Scan(&propertyTax)

	fmt.Println("How much will your MONTHLY hoa fee be?")
	fmt.Scan(&hoa)

	fmt.Println("How much will your MONTHLY Homeowner's Insurance be?")
	fmt.Scan(&homeInsurance)

	fmt.Printf("%.2f", monthlyPayment(housePrice, 4.375, 4422, 82, 75))
}

//Mortgage Calculator, params are HOUSE PRICE, INTEREST RATE, ANNUAL PROPERTY TAX, MONTHLY HOA, MONTHLY HOMEOWNERS INSURANCE
func monthlyPayment(housePrice float64, interestRate float64, propertyTax float64, hoa float64, homeInsurance float64) float64 {
	//house + va funding fee
	initialPrincipal := (housePrice * vaFundingFee) + housePrice

	//converts annual interest rate to decimal figure
	annualInterest := interestRate / 1200

	//Take all data and boils down to a multiplier applied to principal to determine monthly payment
	multiplierDate := math.Pow(1+annualInterest, lengthOfLoan)

	//applies multiplier to loan principal
	multiplierToLoanPrincipal := (multiplierDate * annualInterest) / (multiplierDate - 1)

	monthlyMortgagePayment := initialPrincipal * multiplierToLoanPrincipal

	fullPITI := monthlyMortgagePayment + (propertyTax / 12) + hoa + homeInsurance

	fmt.Println("Monthly Mortgage Payment:")
	fmt.Printf("$%.2f\n", monthlyMortgagePayment)
	fmt.Println("Monthly Payment including Mortgage, Homeowner's Insurance, Property Tax, HOA:")
	fmt.Printf("$%.2f\n", fullPITI)
	return fullPITI
}
