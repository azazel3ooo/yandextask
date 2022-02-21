package models

import (
	"errors"
	"github.com/google/uuid"
)

var (
	Store Storage
)

func init() {
	Store.Init()
}

func InitConfig() Config {
	// изменить на считывание из файла, когда это можно будет
	return Config{Addr: ":8080"}
}

func InitData() Data {
	return make(map[string]string)
}

func (s *Storage) Init() {
	s.Data = InitData()
}

func (s *Storage) Get(key string) (string, error) {
	if v, ok := s.Data[key]; !ok {
		return "", errors.New("неизвестный id")
	} else {
		return v, nil
	}
}

func (s *Storage) Set(val string) string {
	id := uuid.New()
	s.Data[id.String()] = val

	return id.String()
}
