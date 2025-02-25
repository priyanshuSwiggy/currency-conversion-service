package main

import (
	"currency-conversion-service/service"
	"fmt"
)

func main() {
	amount := 100.0
	fromCurrency := service.USD
	toCurrency := service.INR

	convertedAmount := service.ConvertMoney(fromCurrency, toCurrency, amount)
	fmt.Printf("%.2f %s is %.2f %s\n", amount, fromCurrency, convertedAmount, toCurrency)
}
