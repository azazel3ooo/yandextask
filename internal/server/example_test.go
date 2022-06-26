package server

import (
	"github.com/azazel3ooo/yandextask/internal/db"
	"github.com/azazel3ooo/yandextask/internal/models"
	"github.com/gofiber/fiber/v2"
	"log"
	"sync"
)

func ExampleServer_Getter() {
	cfg := models.Config{}
	d := db.Database{}

	err := cfg.Init()
	if err != nil {
		log.Fatal(err)
	}
	d.Init(cfg)

	// chan unused in this handler
	s := NewServer(&d, cfg, fiber.New(), nil)
	s.App.Get("/:id", s.Getter)

	log.Fatal(s.App.Listen(s.Cfg.ServerAddress))
}

func ExampleServer_Setter() {
	cfg := models.Config{}
	d := db.Database{}

	err := cfg.Init()
	if err != nil {
		log.Fatal(err)
	}
	d.Init(cfg)

	// chan unused in this handler
	s := NewServer(&d, cfg, fiber.New(), nil)
	s.App.Post("/", s.Setter)

	log.Fatal(s.App.Listen(s.Cfg.ServerAddress))
}

func ExampleServer_JSONSetter() {
	cfg := models.Config{}
	d := db.Database{}

	err := cfg.Init()
	if err != nil {
		log.Fatal(err)
	}
	d.Init(cfg)

	// chan unused in this handler
	s := NewServer(&d, cfg, fiber.New(), nil)
	s.App.Post("/api/shorten", s.JSONSetter)

	log.Fatal(s.App.Listen(s.Cfg.ServerAddress))
}

func ExampleServer_UserUrlsGet() {
	cfg := models.Config{}
	d := db.Database{}

	err := cfg.Init()
	if err != nil {
		log.Fatal(err)
	}
	d.Init(cfg)

	// chan unused in this handler
	s := NewServer(&d, cfg, fiber.New(), nil)
	s.App.Get("/api/user/urls", s.UserUrlsGet)

	log.Fatal(s.App.Listen(s.Cfg.ServerAddress))
}

func ExampleServer_Ping() {
	cfg := models.Config{}
	d := db.Database{}

	err := cfg.Init()
	if err != nil {
		log.Fatal(err)
	}
	d.Init(cfg)

	// chan unused in this handler
	s := NewServer(&d, cfg, fiber.New(), nil)
	s.App.Get("/api/ping", s.Ping)

	log.Fatal(s.App.Listen(s.Cfg.ServerAddress))
}

func ExampleServer_SetMany() {
	cfg := models.Config{}
	d := db.Database{}

	err := cfg.Init()
	if err != nil {
		log.Fatal(err)
	}
	d.Init(cfg)

	// chan unused in this handler
	s := NewServer(&d, cfg, fiber.New(), nil)
	s.App.Post("/api/shorten/batch", s.SetMany)

	log.Fatal(s.App.Listen(s.Cfg.ServerAddress))
}

func ExampleServer_AsyncDelete() {
	cfg := models.Config{}
	d := db.Database{}
	chanForDelete := make(chan []string)

	err := cfg.Init()
	if err != nil {
		log.Fatal(err)
	}
	d.Init(cfg)
	s := NewServer(&d, cfg, fiber.New(), nil)

	// запускаем вотчер для асинхронного удаления
	wt := sync.WaitGroup{}
	wt.Add(1)
	go FanIn(chanForDelete, &wt, s.Storage)

	s.App.Delete("/api/user/urls", s.AsyncDelete)
	log.Println(s.App.Listen(s.Cfg.ServerAddress))

	// ждем завершения всех удалений
	wt.Wait()
}