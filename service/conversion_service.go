package service

import (
	"currency-conversion-service/money"
	"currency-conversion-service/util"
	"fmt"
)

func ConvertMoney(from money.Money, toCurrency string) (money.Money, error) {
	fromRate, ok := util.ConversionRates[from.Currency]
	if !ok {
		return money.Money{}, fmt.Errorf("invalid from service: %s", from.Currency)
	}
	toRate, ok := util.ConversionRates[toCurrency]
	if !ok {
		return money.Money{}, fmt.Errorf("invalid to service: %s", toCurrency)
	}
	convertedAmount := from.Amount * fromRate / toRate
	return money.Money{Currency: toCurrency, Amount: convertedAmount}, nil
}
