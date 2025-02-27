package service

import (
	"currency-conversion-service/money"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConvertMoney(dsn string, from money.Money, toCurrency string) (money.Money, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return money.Money{}, err
	}

	var fromRate, toRate float64
	if err := db.Table("conversion_rates").Select("rate").Where("currency = ?", from.Currency).Scan(&fromRate).Error; err != nil {
		return money.Money{}, fmt.Errorf("invalid from currency: %s", from.Currency)
	}
	if err := db.Table("conversion_rates").Select("rate").Where("currency = ?", toCurrency).Scan(&toRate).Error; err != nil {
		return money.Money{}, fmt.Errorf("invalid to currency: %s", toCurrency)
	}

	convertedAmount := from.Amount * toRate / fromRate
	return money.Money{Currency: toCurrency, Amount: convertedAmount}, nil
}
