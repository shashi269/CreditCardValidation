package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// check if the credit card number has any four or more consecutive identical digits
func hasFourConsecutiveIdenticalDigits(cardNumber string) bool {
	for i := 0; i < len(cardNumber)-3; i++ {
		if cardNumber[i] == cardNumber[i+1] && cardNumber[i] == cardNumber[i+2] && cardNumber[i] == cardNumber[i+3] {
			return true
		}
	}
	return false
}

// check if the credit card number is valid according to the requirements
func isValidCreditCard(cardNumber string) bool {
	// Regular expression to match
	regexPattern := `^([456]\d{3})(-?\d{4})(-?\d{4})(-?\d{4})$`

	regex := regexp.MustCompile(regexPattern)

	// Check if card number matches the pattern
	if !regex.MatchString(cardNumber) {
		return false
	}

	// Remove hyphens for consecutive digit check
	cardNumberNoHyphens := strings.ReplaceAll(cardNumber, "-", "")

	// Check for four or more consecutive identical digits
	if hasFourConsecutiveIdenticalDigits(cardNumberNoHyphens) {
		return false
	}

	return true
}

func main() {
	// Read the input
	scanner := bufio.NewScanner(os.Stdin)

	// Read the first line for the number of credit cards to validate
	fmt.Println("Enter the number of credit card numbers you want to validate:")
	scanner.Scan()
	nStr := scanner.Text()

	// Parse the string into an integer
	n, err := strconv.Atoi(nStr)
	if err != nil {
		fmt.Println("Error parsing number of credit cards:", err)
		return
	}

	// Read each credit card number and validate them
	for i := 0; i < n; i++ {
		fmt.Println("Enter credit card number", i+1, ":")
		if scanner.Scan() {
			cardNumber := scanner.Text()

			// Validate the credit card number
			if isValidCreditCard(cardNumber) {
				fmt.Println("Valid")
			} else {
				fmt.Println("Invalid")
			}
		}
	}
}
