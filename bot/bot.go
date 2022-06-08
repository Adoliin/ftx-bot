package bot

import (
	"encoding/json"
	"fmt"
	"ftx-bot/models"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"strings"
	"sync"
	"time"

	"gorm.io/gorm"
)

const FTX_API_URL = "https://ftx.com/api"

type BotService struct {
	Mu         sync.Mutex
	db         *gorm.DB
	Markets    []string
	Frequency  time.Duration
	loopsCount uint
}

func Start(db *gorm.DB, botFrequency int, MARKTES string) (*BotService, error) {
	bs := BotService{
		db:         db,
		Markets:    strings.Split(MARKTES, ","),
		Frequency:  time.Duration(botFrequency) * time.Second,
		loopsCount: 0,
	}

	// Check the existence of markets
	for _, market := range bs.Markets {
		ok, err := bs.CheckMarketExistance(market)
		if err != nil {
			return &BotService{}, err
		}
		if !ok {
			return &BotService{}, fmt.Errorf("Market: '%v' does not exist!", market)
		}
	}

	log.Println("🤖 Bot started")
	log.Printf("🤖 Bot - Bot frequency is set at %v\n", bs.Frequency)
	log.Printf("🤖 Bot - The markets that will be fetched are: %v\n", bs.Markets)
	return &bs, nil
}

type MarketTradingVolumeContainer struct {
	mu   sync.Mutex
	list []models.MarketTradingVolume
}

func (bs *BotService) MainLoop() {
	for {
		// marketTradingVolumeListCh := make(chan []models.MarketTradingVolume)
		var marketTradingVolumeContainer MarketTradingVolumeContainer

		bs.loopsCount += 1
		log.Printf("🤖 Bot - Loop n%v\n", bs.loopsCount)

		var wg sync.WaitGroup
		wg.Add(len(bs.Markets))
		for _, market := range bs.Markets {
			log.Printf("🤖 Bot - getting market 24h trading volume for '%v'\n", market)

			//launch each request for the different markets in a goroutine
			go func(market string) {
				defer wg.Done()
				// defer close(marketTradingVolumeListCh)

				change24h, err := bs.GetMarket24hTradingVolume(market)
				if err != nil {
					log.Printf("🤖 Bot - error getting trading volume from '%v' because of: %v\n", market, err)
				}

				marketTradingVolumeContainer.mu.Lock()
				defer marketTradingVolumeContainer.mu.Unlock()
				marketTradingVolumeContainer.list = append(
					marketTradingVolumeContainer.list,
					models.MarketTradingVolume{
						MarketName: market,
						Change24h:  change24h,
					},
				)
			}(market)
			runtime.Gosched()
		}
		wg.Wait()

		bs.db.Create(&marketTradingVolumeContainer.list)
		time.Sleep(bs.Frequency)
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
