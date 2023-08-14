package main

import (
	"fmt"
	"strconv"
)

func isValidCreditCardNumber(cardNumber string) bool {
	cardNumber = cleanCardNumber(cardNumber)
	
	if len(cardNumber) < 13 || len(cardNumber) > 19 {
		return false
	}

	total := 0
	isSecondDigit := false
	for i := len(cardNumber) - 1; i >= 0; i-- {
		digit, err := strconv.Atoi(string(cardNumber[i]))
		if err != nil {
			return false
		}

		if isSecondDigit {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		total += digit
		isSecondDigit = !isSecondDigit
	}

	return total%10 == 0
}

func cleanCardNumber(cardNumber string) string {
	result := ""
	for _, digit := range cardNumber {
		if digit >= '0' && digit <= '9' {
			result += string(digit)
		}
	}
	return result
}

func main() {
	var cardNumber string
	fmt.Print("Enter credit card number: ")
	fmt.Scanln(&cardNumber)

	isValid := isValidCreditCardNumber(cardNumber)
	
	if isValid {
		fmt.Println("Credit card number is valid!")
	} else {
		fmt.Println("Credit card number is invalid!")
	}
}
