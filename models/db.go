package models

import "gorm.io/gorm"

type MarketTradingVolume struct {
	gorm.Model
	Name      string `json:"name"`
	Change24h string `json:"change24h"`
}
