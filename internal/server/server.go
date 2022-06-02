package server

import (
	"github.com/azazel3ooo/yandextask/internal/models"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	Storage       models.Storable
	App           *fiber.App
	Cfg           models.Config
	ChanForDelete chan []string
}

func NewServer(store models.Storable, cfg models.Config, app *fiber.App, c chan []string) (s Server) {
	s.Storage = store
	s.Cfg = cfg
	s.App = app
	s.ChanForDelete = c
	return s
}
