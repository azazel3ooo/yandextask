// Package models - описывает основные структуры и их методы
package models

// Config структура конфигурации сервиса
type Config struct {
	ServerAddress   string `env:"SERVER_ADDRESS"`
	URLBase         string `env:"BASE_URL"`
	FileStoragePath string `env:"FILE_STORAGE_PATH"`
	DatabaseDsn     string `env:"DATABASE_DSN"`
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

type CustomIDSet struct {
	CorrelationID string `json:"correlation_id"`
	OriginalURL   string `json:"original_url,omitempty"`
	ShortURL      string `json:"short_url,omitempty"`
}

type UserResponse struct {
	Short    string `json:"short_url"`
	Original string `json:"original_url"`
}

// Storable интерфейс, которому должно соответствовать хранилище сервиса
type Storable interface {
	userStorable
	urlsStorable
}

// userStorable содержит методы необхожимые для работы с таблицами пользователя
type userStorable interface {
	UsersGet(id string) ([]string, error)
	UsersSet(id, url string) error
	GetUrlsForUser(ids []string) ([]UserResponse, error)
}

// urlsStorable содержит методы необхожимые для работы с таблицами урлов
type urlsStorable interface {
	Get(key string) (string, error)
	Set(val, pth string) (string, error)
	Ping() error
	InsertMany(m []CustomIDSet) ([]CustomIDSet, error)
	Delete(ids []string) error
}
