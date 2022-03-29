package service

import (
	"github.com/azazel3ooo/yandextask/internal/app/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gorilla/securecookie"
	"log"
)

func StartService() {
	var (
		cfg   models.Config
		store models.Storable
	)

	err := cfg.Init()
	if err != nil {
		log.Fatal(err.Error())
	}
	if cfg.DatabaseDsn == "" {
		var s models.Storage
		s.Init(cfg)
		store = &s
	} else {
		var s models.Database
		s.Init(cfg)
		store = &s
	}

	app := fiber.New()
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))
	app.Use(recover.New(recover.Config{EnableStackTrace: true}))
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${latency} ${method} ${path} ${resBody}\n",
	}))

	s := models.NewServer(store, cfg, app)

	hashKey := []byte("top-secret")
	blockKey := []byte("super-secret")
	s.Cookie = securecookie.New(hashKey, blockKey)

	s.App.Get("/:id", s.Getter)
	s.App.Post("/", s.Setter)
	s.App.Post("/api/shorten", s.JSONSetter)
	s.App.Get("/api/user/urls", s.UserUrlsGet)
	s.App.Get("/ping", s.Ping)
	s.App.Post("/api/shorten/batch", s.SetMany)
	log.Fatal(s.App.Listen(s.Cfg.ServerAddress))
}
