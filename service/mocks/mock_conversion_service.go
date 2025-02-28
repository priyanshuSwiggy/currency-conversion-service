package mocks

import (
	"currency-conversion-service/money"
	"github.com/stretchr/testify/mock"
)

type MockDB struct {
	mock.Mock
}

func (m *MockDB) GetRate(currency string) (float64, error) {
	args := m.Called(currency)
	return args.Get(0).(float64), args.Error(1)
}

func ConvertMoneyWithMock(db *MockDB, from money.Money, toCurrency string) (money.Money, error) {
	fromRate, err := db.GetRate(from.Currency)
	if err != nil {
		return money.Money{}, err
	}

	toRate, err := db.GetRate(toCurrency)
	if err != nil {
		return money.Money{}, err
	}

	convertedAmount := from.Amount * toRate / fromRate
	return money.Money{Currency: toCurrency, Amount: convertedAmount}, nil
}
