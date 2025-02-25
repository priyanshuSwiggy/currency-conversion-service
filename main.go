package main

import (
	"currency-conversion-service/convert"
	"fmt"
)

func main() {
	amount := 100.0
	fromCurrency := convert.USD
	toCurrency := convert.INR

	convertedAmount := convert.Convert(fromCurrency, toCurrency, amount)
	fmt.Printf("%.2f %s is %.2f %s\n", amount, fromCurrency, convertedAmount, toCurrency)
}
