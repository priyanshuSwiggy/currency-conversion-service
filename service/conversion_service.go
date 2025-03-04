package service

import (
	"currency-conversion-service/dao"
	"currency-conversion-service/money"
	"fmt"
	"log"
)

func ConvertMoney(from money.Money, toCurrency string) (money.Money, error) {
	log.Printf("Fetching exchange rate for: %s", from.Currency)

	fromRate, err := dao.GetRate(from.Currency)
	if err != nil {
		return money.Money{}, fmt.Errorf("failed to fetch rate for currency %s: %v", from.Currency, err)
	}

	log.Printf("Fetching exchange rate for: %s", toCurrency)

	toRate, err := dao.GetRate(toCurrency)
	if err != nil {
		return money.Money{}, fmt.Errorf("failed to fetch rate for currency %s: %v", toCurrency, err)
	}

	if fromRate == 0 {
		return money.Money{}, fmt.Errorf("invalid fromRate: %f for currency %s", fromRate, from.Currency)
	}

	if toRate == 0 {
		return money.Money{}, fmt.Errorf("invalid toRate: %f for currency %s", toRate, toCurrency)
	}

	convertedAmount := (from.Amount / fromRate) * toRate
	log.Printf("Converted %.2f %s to %.2f %s", from.Amount, from.Currency, convertedAmount, toCurrency)

	return money.Money{Currency: toCurrency, Amount: convertedAmount}, nil
}
