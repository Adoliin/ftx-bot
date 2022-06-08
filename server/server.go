package server

import (
	"fmt"
	"ftx-bot/bot"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"
)

type Server struct {
	db *gorm.DB
	bs *bot.BotService
}

func Start(db *gorm.DB, bs *bot.BotService) {
	s := Server{db: db, bs: bs}
	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Format: "💻 API - [${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	app.Get("/bot/frequency", s.getBotFrequency)
	app.Put("/bot/frequency/:new_frequency", s.changeBotFrequency)
	app.Listen(":3000")
}

type GetBotFrequenctResponse struct {
	Bot_frequency float64 `json:"bot_frequency_seconds"`
}

func (s *Server) getBotFrequency(c *fiber.Ctx) error {
	return c.JSON(GetBotFrequenctResponse{
		Bot_frequency: s.bs.Frequency.Seconds(),
	})
}

func (s *Server) changeBotFrequency(c *fiber.Ctx) error {
	new_frequency, err := strconv.Atoi(c.Params("new_frequency"))
	if err != nil {
		fmt.Println(err)
		return c.Status(400).JSON("Invalid input")
	}

	s.bs.Mu.Lock()
	defer s.bs.Mu.Unlock()
	s.bs.Frequency = time.Duration(new_frequency) * time.Second

	c.Status(200)

	return nil
}
