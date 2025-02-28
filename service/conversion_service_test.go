package service

import (
	"currency-conversion-service/money"
	"currency-conversion-service/service/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertUSDToINR(t *testing.T) {
	mockDB := new(mocks.MockDB)
	mockDB.On("GetRate", "TESTUSD").Return(1.0, nil)
	mockDB.On("GetRate", "TESTINR").Return(87.191146, nil)

	from := money.Money{Currency: "TESTUSD", Amount: 100.0}
	toCurrency := "TESTINR"
	expectedAmount := 8719.1146

	converted, err := mocks.ConvertMoneyWithMock(mockDB, from, toCurrency)
	assert.NoError(t, err)
	assert.Equal(t, expectedAmount, converted.Amount)

	mockDB.AssertExpectations(t)
}

func TestConvertINRToUSD(t *testing.T) {
	mockDB := new(mocks.MockDB)
	mockDB.On("GetRate", "TESTINR").Return(87.191146, nil)
	mockDB.On("GetRate", "TESTUSD").Return(1.0, nil)

	from := money.Money{Currency: "TESTINR", Amount: 8719.1146}
	toCurrency := "TESTUSD"
	expectedAmount := 100.0

	converted, err := mocks.ConvertMoneyWithMock(mockDB, from, toCurrency)
	assert.NoError(t, err)
	assert.Equal(t, expectedAmount, converted.Amount)

	mockDB.AssertExpectations(t)
}

func TestConvertEURToINR(t *testing.T) {
	mockDB := new(mocks.MockDB)
	mockDB.On("GetRate", "TESTEUR").Return(0.954039, nil)
	mockDB.On("GetRate", "TESTINR").Return(87.191146, nil)

	from := money.Money{Currency: "TESTEUR", Amount: 90.0}
	toCurrency := "TESTINR"
	expectedAmount := 8225.2435

	converted, err := mocks.ConvertMoneyWithMock(mockDB, from, toCurrency)
	assert.NoError(t, err)
	assert.InEpsilon(t, expectedAmount, converted.Amount, 0.01)

	mockDB.AssertExpectations(t)
}

func TestConvertINRToEUR(t *testing.T) {
	mockDB := new(mocks.MockDB)
	mockDB.On("GetRate", "TESTINR").Return(87.191146, nil)
	mockDB.On("GetRate", "TESTEUR").Return(0.954039, nil)

	from := money.Money{Currency: "TESTINR", Amount: 8227.20314}
	toCurrency := "TESTEUR"
	expectedAmount := 90.021

	converted, err := mocks.ConvertMoneyWithMock(mockDB, from, toCurrency)
	assert.NoError(t, err)
	assert.InEpsilon(t, expectedAmount, converted.Amount, 0.01)

	mockDB.AssertExpectations(t)
}

func TestConvertUSDToEUR(t *testing.T) {
	mockDB := new(mocks.MockDB)
	mockDB.On("GetRate", "TESTUSD").Return(1.0, nil)
	mockDB.On("GetRate", "TESTEUR").Return(0.954039, nil)

	from := money.Money{Currency: "TESTUSD", Amount: 100.0}
	toCurrency := "TESTEUR"
	expectedAmount := 95.4039

	converted, err := mocks.ConvertMoneyWithMock(mockDB, from, toCurrency)
	assert.NoError(t, err)
	assert.InEpsilon(t, expectedAmount, converted.Amount, 0.01)

	mockDB.AssertExpectations(t)
}

func TestConvertEURToUSD(t *testing.T) {
	mockDB := new(mocks.MockDB)
	mockDB.On("GetRate", "TESTEUR").Return(0.954039, nil)
	mockDB.On("GetRate", "TESTUSD").Return(1.0, nil)

	from := money.Money{Currency: "TESTEUR", Amount: 90.0}
	toCurrency := "TESTUSD"
	expectedAmount := 94.312

	converted, err := mocks.ConvertMoneyWithMock(mockDB, from, toCurrency)
	assert.NoError(t, err)
	assert.InEpsilon(t, expectedAmount, converted.Amount, 0.01)

	mockDB.AssertExpectations(t)
}
