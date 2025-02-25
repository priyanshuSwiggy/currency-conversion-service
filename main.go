package main

import (
	"currency-conversion-service/currency"
	"fmt"
)

func main() {
	amount := 100.0
	fromCurrency := currency.USD
	toCurrency := currency.INR

	convertedAmount := currency.Convert(fromCurrency, toCurrency, amount)
	fmt.Printf("%.2f %s is %.2f %s\n", amount, fromCurrency, convertedAmount, toCurrency)
}
