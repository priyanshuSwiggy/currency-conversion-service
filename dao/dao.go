package dao

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ExchangeRate struct {
	Currency string  `json:"currency"`
	Rate     float64 `json:"rate"`
}

func ConnectDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func UpdateRateInDB(db *gorm.DB, rate ExchangeRate) error {
	return db.Table("conversion_rates").Where("currency = ?", rate.Currency).Updates(map[string]interface{}{"rate": rate.Rate}).Error
}
