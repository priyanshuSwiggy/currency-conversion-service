package main

import (
	"currency-conversion-service/util"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsn := "host=localhost user=root password=root dbname=conversiondb port=5432 sslmode=disable"
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	if err := util.LoadConversionRates(dsn); err != nil {
		log.Fatal("Failed to load conversion rates:", err)
	}
}
