package models

import "github.com/gofiber/fiber/v2"

type Server struct {
	Storage Storable
	App     *fiber.App
	Cfg     Config
}

type Config struct {
	ServerAddress   string `env:"SERVER_ADDRESS"`
	URLBase         string `env:"BASE_URL"`
	FileStoragePath string `env:"FILE_STORAGE_PATH"`
}

type Data map[string]string

type Storage struct {
	Data Data
}

type Request struct {
	Addr string `json:"url"`
}

type Response struct {
	Result string `json:"result"`
}

type Storable interface {
	Get(key string) (string, error)
	Set(val, pth string) string
}
