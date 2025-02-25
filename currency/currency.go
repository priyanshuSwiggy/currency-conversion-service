package currency

import "fmt"

type Currency int

const (
	INR Currency = iota
	USD
	EUR
)

var conversionRates = map[Currency]float64{
	INR: 1.0,
	USD: 83.0,
	EUR: 90.0,
}

var currencyCodes = map[string]Currency{
	"INR": INR,
	"USD": USD,
	"EUR": EUR,
}

func (c Currency) ConversionRate() float64 {
	return conversionRates[c]
}

func Convert(fromCurrency, toCurrency Currency, amount float64) float64 {
	return amount * fromCurrency.ConversionRate() / toCurrency.ConversionRate()
}

func GetCurrencyByCode(code string) (Currency, error) {
	currency, exists := currencyCodes[code]
	if !exists {
		return 0, fmt.Errorf("invalid currency code: %s", code)
	}
	return currency, nil
}
