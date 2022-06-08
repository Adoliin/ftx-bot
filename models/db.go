package models

import "gorm.io/gorm"

type MarketTradingVolume struct {
	gorm.Model
	MarketName string  `json:"market_name"`
	Change24h  float64 `json:"change24h"`
}
