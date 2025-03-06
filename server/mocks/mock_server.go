package mocks

import (
	"currency-conversion-service/money"
	"fmt"
)

type MockConverter struct{}

func (m *MockConverter) ConvertMoney(from money.Money, toCurrency string) (money.Money, error) {
	if toCurrency == "XYZ" {
		return money.Money{}, fmt.Errorf("unsupported currency")
	}
	return money.Money{Currency: toCurrency, Amount: 85}, nil
}
