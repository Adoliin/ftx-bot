package models

type FtxMarket struct {
	Success bool `json:"success"`
	Result  struct {
		Name                  string      `json:"name"`
		Enabled               bool        `json:"enabled"`
		PostOnly              bool        `json:"postOnly"`
		PriceIncrement        float64     `json:"priceIncrement"`
		SizeIncrement         float64     `json:"sizeIncrement"`
		MinProvideSize        float64     `json:"minProvideSize"`
		Last                  float64     `json:"last"`
		Bid                   float64     `json:"bid"`
		Ask                   float64     `json:"ask"`
		Price                 float64     `json:"price"`
		Type                  string      `json:"type"`
		BaseCurrency          string      `json:"baseCurrency"`
		IsEtfMarket           bool        `json:"isEtfMarket"`
		QuoteCurrency         string      `json:"quoteCurrency"`
		Underlying            interface{} `json:"underlying"`
		Restricted            bool        `json:"restricted"`
		HighLeverageFeeExempt bool        `json:"highLeverageFeeExempt"`
		LargeOrderThreshold   float64     `json:"largeOrderThreshold"`
		Change1H              float64     `json:"change1h"`
		Change24H             float64     `json:"change24h"`
		ChangeBod             float64     `json:"changeBod"`
		QuoteVolume24H        float64     `json:"quoteVolume24h"`
		VolumeUsd24H          float64     `json:"volumeUsd24h"`
		PriceHigh24H          float64     `json:"priceHigh24h"`
		PriceLow24H           float64     `json:"priceLow24h"`
	} `json:"result"`
}
