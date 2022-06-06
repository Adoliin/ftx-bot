package bot

import (
	"encoding/json"
	"fmt"
	"ftx-bot/models"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"gorm.io/gorm"
)

const FTX_API_URL = "https://ftx.com/api"

type BotService struct {
	db         *gorm.DB
	markets    []string
	frequency  time.Duration
	loopsCount uint
}

func Start(db *gorm.DB, botFrequency int, MARKTES string) (BotService, error) {
	bs := BotService{
		db:         db,
		markets:    strings.Split(MARKTES, ","),
		frequency:  time.Duration(botFrequency) * time.Second,
		loopsCount: 0,
	}

	// Check the existence of markets
	for _, market := range bs.markets {
		ok, err := bs.CheckMarketExistance(market)
		if err != nil {
			return BotService{}, err
		}
		if !ok {
			return BotService{}, fmt.Errorf("Market: '%v' does not exist!", market)
		}
	}

	log.Println("Bot started")
	return bs, nil
}

func (bs *BotService) MainLoop() {
	for {
		bs.loopsCount += 1
		log.Printf("Bot - Starting loop n%v\n", bs.loopsCount)
		for _, market := range bs.markets {
			log.Printf("Bot - getting market 24h trading volume for '%v'\n", market)

			// do a go routine here
			change24h, err := bs.GetMarket24hTradingVolume(market)
			log.Printf("change24h -> %v\n", change24h)
			if err != nil {
				log.Printf("Bot - error getting trading volume from '%v'\n", market)
			}
			//insert shit in database here
		}

		time.Sleep(bs.frequency)
	}
}

func (bs *BotService) GetMarket24hTradingVolume(market string) (float64, error) {
	resp, err := http.Get(
		fmt.Sprintf("%s/%s/%s", FTX_API_URL, "markets", market),
	)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var res models.FtxMarket
	if err := json.Unmarshal(body, &res); err != nil {
		return 0, err
	}

	return res.Result.Change24H, nil
}

func (bs *BotService) CheckMarketExistance(market string) (bool, error) {
	resp, err := http.Get(
		fmt.Sprintf("%s/%s/%s", FTX_API_URL, "markets", market),
	)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return false, nil
	}
	return true, nil
}
