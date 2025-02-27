package util

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func UpdateRatesInDB(dsn string, rates map[string]float64) error {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	for currency, rate := range rates {
		if err := db.Table("conversion_rates").Where("currency = ?", currency).Updates(map[string]interface{}{"rate": rate}).Error; err != nil {
			return err
		}
	}

	return nil
}
