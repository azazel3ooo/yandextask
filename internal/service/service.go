// Package service - запускат сервис
package service

import (
	"github.com/azazel3ooo/yandextask/internal/logic"
	"github.com/azazel3ooo/yandextask/internal/transport/grpc"
	"github.com/azazel3ooo/yandextask/internal/transport/http_transport"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"sync"

	"github.com/azazel3ooo/yandextask/internal/db"
	"github.com/azazel3ooo/yandextask/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// StartService производит инициализацию основных струтур с помощью их методов и запускает сервис
func StartService() {
	var (
		cfg   models.Config
		store models.Storable
	)

	err := cfg.Init()
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println(cfg.DatabaseDsn)
	if cfg.DatabaseDsn == "" {
		var s models.Storage
		s.Init(cfg)
		store = &s
	} else {
		var s db.Database
		s.Init(cfg)
		store = &s
		defer s.Conn.Close()
	}

	app := fiber.New()
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))
	app.Use(recover.New(recover.Config{EnableStackTrace: true}))
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${latency} ${method} ${path} ${resBody}\n",
	}))
	app.Use(pprof.New())

	chanForDelete := make(chan []string, 10)
	s := http_transport.NewServer(store, cfg, app, chanForDelete)
	wt := sync.WaitGroup{}
	wt.Add(1)
	go logic.FanIn(chanForDelete, &wt, s.Storage)

	s.App.Get("/ping", s.Ping)
	s.App.Get("/:id", s.Getter)
	s.App.Post("/", s.Setter)
	s.App.Post("/api/shorten", s.JSONSetter)
	s.App.Get("/api/user/urls", s.UserUrlsGet)
	s.App.Post("/api/shorten/batch", s.SetMany)
	s.App.Delete("/api/user/urls", s.AsyncDelete)
	s.App.Get("/api/internal/stats", s.GetStat)

	connectionsClosed := make(chan struct{})
	go s.GracefulShutdown(connectionsClosed)

	if err = s.Listen(); err == http.ErrServerClosed {
		log.Println(err)
	} else {
		<-connectionsClosed
		log.Println("Success closed")
	}

	close(chanForDelete)
	wt.Wait()
}

// Нужно бы объединить в одну функцию с StartService, но из описания задания не было понятно, как они должны функционировать вместе
// и запускаться

func StartGRPC() {
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
		var s db.Database
		s.Init(cfg)
		store = &s
		defer s.Conn.Close()
	}

	chanForDelete := make(chan []string, 10)
	s, err := grpc.GetServer(cfg, chanForDelete, store)
	if err != nil {
		log.Fatal(err)
	}

	wt := sync.WaitGroup{}
	wt.Add(1)
	go logic.FanIn(chanForDelete, &wt, store)
	close(chanForDelete)

	listen, err := net.Listen("tcp", cfg.GAddress)
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(listen); err != nil {
		log.Fatal(err)
	}

	wt.Wait()
}
