package convert

import (
	"currency-conversion-service/config"
	"currency-conversion-service/money"
	"fmt"
)

func Convert(from money.Money, toCurrency string) (money.Money, error) {
	fromRate, ok := config.ConversionRates[from.Currency]
	if !ok {
		return money.Money{}, fmt.Errorf("invalid from convert: %s", from.Currency)
	}
	toRate, ok := config.ConversionRates[toCurrency]
	if !ok {
		return money.Money{}, fmt.Errorf("invalid to convert: %s", toCurrency)
	}
	convertedAmount := from.Amount * fromRate / toRate
	return money.Money{Currency: toCurrency, Amount: convertedAmount}, nil
}
