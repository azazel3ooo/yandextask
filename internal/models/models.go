// Package models - описывает основные структуры и их методы
package models

// Config структура конфигурации сервиса
type Config struct {
	ServerAddress   string `env:"SERVER_ADDRESS"`
	URLBase         string `env:"BASE_URL"`
	FileStoragePath string `env:"FILE_STORAGE_PATH"`
	DatabaseDsn     string `env:"DATABASE_DSN"`
	EnableTLS       bool   `env:"ENABLE_HTTPS"`
}

type data map[string]string

type users map[string][]string

// Storage структура для хранения данных в памяти
type Storage struct {
	Data  data
	Users users
}

// Request - структура запроса для добавления ссылки
type Request struct {
	Addr string `json:"url"`
}

// Response - структура ответа после добавления короткой ссылки
type Response struct {
	Result string `json:"result"`
}

// CustomIDSet - структура запроса\ответа для множественного добавления коротких ссылок
type CustomIDSet struct {
	CorrelationID string `json:"correlation_id"`
	OriginalURL   string `json:"original_url,omitempty"`
	ShortURL      string `json:"short_url,omitempty"`
}

// UserResponse - структура для вывода ссылок пользователя
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
