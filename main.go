package main

import (
	"fmt"
	"math"
)

const vaFundingFee float64 = 0.0215
const lengthOfLoan float64 = 360

func main() {
	fmt.Printf("%.2f", monthlyPayment(450000, 4.375, 4422, 82, 75))

}

//Mortgage Calculator, params are HOUSE PRICE, INTEREST RATE, ANNUAL PROPERTY TAX, MONTHLY HOA, MONTHLY HOMEOWNERS INSURANCE
func monthlyPayment(housePrice float64, interestRate float64, propertyTax float64, hoa float64, homeInsurance float64) float64 {
	//house + va funding fee
	initialPrincipal := (housePrice * vaFundingFee) + housePrice

	//converts annual interest rate to decimal figure
	annualInterest := interestRate / 1200

	//Take all data and boils down to a multiplier applied to principal to determine monthly payment
	multiplierDate := math.Pow((1 + annualInterest), lengthOfLoan)

	//applies multiplier to loan principal
	multiplierToLoanPrincipal := (multiplierDate * annualInterest) / (multiplierDate - 1)

	monthlyMortgagePayment := initialPrincipal * multiplierToLoanPrincipal

	fullPITI := monthlyMortgagePayment + (propertyTax / 12) + hoa + homeInsurance

	fmt.Printf("Monthly Mortgage Payment => \t\t\t\t\t\t\t\t\t\t\t\t\t\t$%.2f\n", monthlyMortgagePayment)
	fmt.Printf("Monthly Payment including Mortgage, Homeowner's Insurance, Property Tax, HOA => \t$%.2f\n", fullPITI)
	return fullPITI
}
