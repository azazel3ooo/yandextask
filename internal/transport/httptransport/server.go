// Package httptransport - содержит описание сруктуры сервера и хендлеров
package httptransport

import (
	"errors"
	"github.com/azazel3ooo/yandextask/internal/models"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

// Server хранит структуры необходимые для работы сервиса
type Server struct {
	Storage       models.Storable
	App           *fiber.App
	Cfg           models.Config
	ChanForDelete chan []string
}

// NewServer возвращает Server с заданыыми параметрами
func NewServer(store models.Storable, cfg models.Config, app *fiber.App, c chan []string) (s Server) {
	s.Storage = store
	s.Cfg = cfg
	s.App = app
	s.ChanForDelete = c
	return s
}

// Listen запускает соответсвующую конфигурацию для http\https
func (s *Server) Listen() (err error) {
	switch s.Cfg.EnableTLS {
	case true:
		cert := filepath.Join(".", "config", "cert.pem")
		key := filepath.Join(".", "config", "cert.key")
		err = s.App.ListenTLS(s.Cfg.ServerAddress, cert, key)

	case false:
		err = s.App.Listen(s.Cfg.ServerAddress)

	default:
		err = errors.New("httptransport.Listen() error: cfg.EnableTLS doesn't have value")
	}
	return err
}

// GracefulShutdown - перехватывает syscall'ы и выполняет штатную остановку приложения
func (s *Server) GracefulShutdown(connectionsClosed chan struct{}) {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	<-sigint

	if err := s.App.Shutdown(); err != nil {
		log.Println(err)
	}
	connectionsClosed <- struct{}{}
}
