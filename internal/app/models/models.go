package models

type Config struct {
	Addr string `yaml:"address"`
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
