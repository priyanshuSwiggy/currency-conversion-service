package service

import (
	"currency-conversion-service/money"
	"currency-conversion-service/util"
	"testing"
)

func TestConvertMoney(t *testing.T) {
	util.ConversionRates = map[string]float64{
		"INR": 1.0,
		"USD": 83.0,
		"EUR": 90.0,
	}

	from := money.Money{Currency: "USD", Amount: 100.0}
	toCurrency := "INR"
	expectedAmount := 8300.0

	converted, err := ConvertMoney(from, toCurrency)
	if err != nil {
		t.Fatalf("Conversion failed: %v", err)
	}

	if converted.Amount != expectedAmount {
		t.Errorf("Expected %v, got %v", expectedAmount, converted.Amount)
	}
}
