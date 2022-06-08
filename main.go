package main

import (
	"log"
	"os"
	"runtime"
	"strconv"
	"sync"

	"ftx-bot/bot"
	"ftx-bot/database"
	"ftx-bot/server"
)

var (
	// Bot
	BOT_MARKETS        = GetEnv("BOT_MARKETS")
	BOT_FREQUENCY, err = strconv.Atoi(GetEnv("BOT_FREQUENCY"))

	// Database
	DB_HOST = GetEnv("DB_HOST")
	DB_PORT = GetEnv("DB_PORT")
	DB_USER = GetEnv("DB_USER")
	DB_PWD  = GetEnv("DB_PWD")
	DB_NAME = GetEnv("DB_NAME")
)

func main() {
	db, err := database.ConnectDb(DB_HOST, DB_PORT, DB_USER, DB_PWD, DB_NAME)
	if err != nil {
		log.Fatalf("Failed to connect to db: %v", err)
	}

	bs, err := bot.Start(db, BOT_FREQUENCY, BOT_MARKETS)
	if err != nil {
		log.Fatalf("Failed to start bot: %v", err)
	}

	var wg sync.WaitGroup
	wg.Add(2)

	// Bot main loop
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		bs.MainLoop()
	}(&wg)
	runtime.Gosched()

	// API Server
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		server.Start(db, bs)
	}(&wg)
	runtime.Gosched()

	wg.Wait()
}

func GetEnv(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		log.Fatal("Environment variable not found: ", key)
	}
	return value
}
