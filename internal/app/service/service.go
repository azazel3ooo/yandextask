package service

import (
	"fmt"
	"github.com/azazel3ooo/yandextask/internal/app/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log"
)

func StartService() {
	var (
		store models.Storage
		cfg   models.Config
	)

	err := cfg.Init()
	if err != nil {
		log.Fatal(err.Error())
	}
	store.Init(cfg)

	app := fiber.New()
	app.Use(recover.New(recover.Config{EnableStackTrace: true}))
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${latency} ${method} ${path} ${resBody}\n",
	}))

	s := models.NewServer(&store, cfg, app)

	s.App.Get("/:id", s.Getter)
	s.App.Post("/", s.Setter)
	s.App.Post("/api/shorten", s.JSONSetter)
	log.Fatal(s.App.Listen(fmt.Sprintf(":%d", s.Cfg.ServerAddress)))
}
