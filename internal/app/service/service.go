package service

import (
	"flag"
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

	flag.StringVar(&cfg.ServerAddress, "a", "8080", "Server address")
	flag.StringVar(&cfg.URLBase, "b", "127.0.0.1", "Base url")
	flag.StringVar(&cfg.FileStoragePath, "c", "./tmp/tmp.txt", "Filepath for backup")
	flag.Parse()

	err := cfg.Init()
	if err != nil {
		log.Fatal(err.Error())
	}
	store.Init(cfg)
	log.Println(cfg.FileStoragePath)
	if cfg.FileStoragePath != "" {
		defer store.Backup(cfg)
	}

	app := fiber.New()
	app.Use(recover.New(recover.Config{EnableStackTrace: true}))
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${latency} ${method} ${path} ${resBody}\n",
	}))

	s := models.NewServer(&store, cfg, app)

	s.App.Get("/:id", s.Getter)
	s.App.Post("/", s.Setter)
	s.App.Post("/api/shorten", s.JSONSetter)
	log.Fatal(s.App.Listen(":" + s.Cfg.ServerAddress))
}
