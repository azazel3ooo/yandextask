package models

import (
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	Storage Storable
	App     *fiber.App
	Cfg     Config
}

type Config struct {
	ServerAddress   string `env:"SERVER_ADDRESS"`
	URLBase         string `env:"BASE_URL"`
	FileStoragePath string `env:"FILE_STORAGE_PATH"`
	DatabaseDsn     string `env:"DATABASE_DSN"`
}

type Database struct {
	name           string
	connectionInfo string
}

type Data map[string]string

type Users map[string][]string

type Storage struct {
	Data  Data
	Users Users
}

type Request struct {
	Addr string `json:"url"`
}

type Response struct {
	Result string `json:"result"`
}

type CustomIdSet struct {
	CorrelationId string `json:"correlation_id"`
	OriginalUrl   string `json:"original_url,omitempty"`
	ShortUrl      string `json:"short_url,omitempty"`
}

type UserResponse struct {
	Short    string `json:"short_url"`
	Original string `json:"original_url"`
}

type Storable interface {
	Get(key string) (string, error)
	Set(val, pth string) (string, error)
	UsersGet(id string) ([]string, error)
	UsersSet(id, url string) error
	GetUrlsForUser(ids []string) ([]UserResponse, error)
	Ping() error
	InsertMany(m []CustomIdSet) ([]CustomIdSet, error)
}
