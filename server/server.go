package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"
)

type Server struct {
	db *gorm.DB
}

func Start(db *gorm.DB) {
	s := Server{db: db}
	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	app.Get("/api/test", s.test)
	app.Listen(":3000")
}

func (s *Server) test(c *fiber.Ctx) error {

	return c.JSON("testing...")
}
