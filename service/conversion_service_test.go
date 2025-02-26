package service

import (
	"currency-conversion-service/money"
	"currency-conversion-service/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertUSDToINR(t *testing.T) {
	util.ConversionRates = map[string]float64{
		"INR": 1.0,
		"USD": 83.0,
	}

	from := money.Money{Currency: "USD", Amount: 100.0}
	toCurrency := "INR"
	expectedAmount := 8300.0

	converted, err := ConvertMoney(from, toCurrency)
	assert.NoError(t, err)
	assert.Equal(t, expectedAmount, converted.Amount)
}

func TestConvertINRToUSD(t *testing.T) {
	util.ConversionRates = map[string]float64{
		"INR": 1.0,
		"USD": 83.0,
	}

	from := money.Money{Currency: "INR", Amount: 8300.0}
	toCurrency := "USD"
	expectedAmount := 100.0

	converted, err := ConvertMoney(from, toCurrency)
	assert.NoError(t, err)
	assert.Equal(t, expectedAmount, converted.Amount)
}

func TestConvertEURToINR(t *testing.T) {
	util.ConversionRates = map[string]float64{
		"INR": 1.0,
		"EUR": 90.0,
	}

	from := money.Money{Currency: "EUR", Amount: 90.0}
	toCurrency := "INR"
	expectedAmount := 8100.0

	converted, err := ConvertMoney(from, toCurrency)
	assert.NoError(t, err)
	assert.Equal(t, expectedAmount, converted.Amount)
}

func TestConvertINRToEUR(t *testing.T) {
	util.ConversionRates = map[string]float64{
		"INR": 1.0,
		"EUR": 90.0,
	}

	from := money.Money{Currency: "INR", Amount: 8100.0}
	toCurrency := "EUR"
	expectedAmount := 90.0

	converted, err := ConvertMoney(from, toCurrency)
	assert.NoError(t, err)
	assert.Equal(t, expectedAmount, converted.Amount)
}

func TestConvertUSDToEUR(t *testing.T) {
	util.ConversionRates = map[string]float64{
		"USD": 83.0,
		"EUR": 90.0,
	}

	from := money.Money{Currency: "USD", Amount: 100.0}
	toCurrency := "EUR"
	expectedAmount := 92.22

	converted, err := ConvertMoney(from, toCurrency)
	assert.NoError(t, err)
	assert.InEpsilon(t, expectedAmount, converted.Amount, 0.01)
}

func TestConvertEURToUSD(t *testing.T) {
	util.ConversionRates = map[string]float64{
		"USD": 83.0,
		"EUR": 90.0,
	}

	from := money.Money{Currency: "EUR", Amount: 90.0}
	toCurrency := "USD"
	expectedAmount := 97.59

	converted, err := ConvertMoney(from, toCurrency)
	assert.NoError(t, err)
	assert.InEpsilon(t, expectedAmount, converted.Amount, 0.01)
}
