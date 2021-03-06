package database

import (
	"fmt"

	"ftx-bot/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDb(
	dbHost string,
	dbPort string,
	dbUser string,
	dbPwd string,
	dbName string,
) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		`host=%v port=%v user=%v password=%v dbname=%v sslmode=disable`,
		dbHost, dbPort, dbUser, dbPwd, dbName,
	)
	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return &gorm.DB{}, err
	}

	db.AutoMigrate(&models.MarketTradingVolume{})

	return db, nil
}
