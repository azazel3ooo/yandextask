package models

type Config struct {
	Addr string `yaml:"address"`
}

type Data map[string]string

type Storage struct {
	Data Data
}
