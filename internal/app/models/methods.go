package models

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"log"
	"os"
	"strings"
)

func (c *Config) Init() error {
	err := env.Parse(c)
	if err != nil {
		return err
	}
	if c.ServerAddress == "" {
		flag.StringVar(&c.ServerAddress, "a", "localhost:8080", "Server address")
	}
	if c.URLBase == "" {
		flag.StringVar(&c.URLBase, "b", "http://127.0.0.1:8080", "Base url")
	}
	if c.FileStoragePath == "" {
		flag.StringVar(&c.FileStoragePath, "f", "./tmp/tmp.txt", "Filepath for backup")
	}
	if c.DatabaseDsn == "" {
		flag.StringVar(&c.DatabaseDsn, "d", "", "Database address")
	}

	flag.Parse()

	return nil
}

func NewServer(store Storable, cfg Config, app *fiber.App) (s Server) {
	s.Storage = store
	s.Cfg = cfg
	s.App = app
	return s
}

func InitData() Data {
	return make(map[string]string)
}

func InitUsers() Users {
	return make(map[string][]string)
}

func UploadData(s *Storage, cfg Config) {
	f, err := os.OpenFile(cfg.FileStoragePath, os.O_RDONLY, 0666)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer f.Close()

	m := make(map[string]string)
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		pair := strings.Split(sc.Text(), ";")
		if len(pair) == 2 {
			m[pair[0]] = pair[1]
		}
	}

	s.Data = m
}

func (s *Storage) Init(cfg Config) {
	s.Data = InitData()
	if cfg.FileStoragePath != "" {
		if _, err := os.Stat(cfg.FileStoragePath); !os.IsNotExist(err) {
			UploadData(s, cfg)
		}
	}
	s.Users = InitUsers()
}

// Get Return original URL
func (s *Storage) Get(key string) (string, error) {
	if v, ok := s.Data[key]; !ok {
		return "", errors.New("неизвестный id")
	} else {
		return v, nil
	}
}

// Set Return URL id in storage
func (s *Storage) Set(val, pth string) (string, error) {
	id := uuid.New()
	s.Data[id.String()] = val

	if pth != "" {
		SetToFile(id.String(), val, pth)
	}

	return id.String(), nil
}

func SetToFile(k, v, pth string) {
	d := []byte(fmt.Sprintf("%s;%s\n", k, v))
	f, err := os.OpenFile(pth, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Println(err.Error())
	}
	defer f.Close()

	_, err = f.Write(d)
	if err != nil {
		log.Println(err.Error())
	}
}

func (s *Storage) UsersGet(id string) ([]string, error) {
	urls, ok := s.Users[id]
	if !ok {
		return nil, errors.New("unknown user")
	}

	return urls, nil
}

func (s *Storage) UsersSet(id, url string) error {
	urls, ok := s.Users[id]
	if !ok {
		s.Users[id] = []string{url}
	} else {
		s.Users[id] = append(urls, url)
	}
	return nil
}

func (s *Storage) GetUrlsForUser(ids []string) ([]UserResponse, error) {
	var res []UserResponse
	for _, id := range ids {
		url, ok := s.Data[id]
		if ok {
			res = append(res, UserResponse{id, url})
		}
	}
	return res, nil
}

func (s *Storage) Ping() error {
	return errors.New("i'm not db")
}

func (s *Storage) InsertMany(m []CustomIdSet) ([]CustomIdSet, error) {
	var res []CustomIdSet
	for _, el := range m {
		s.Data[el.CorrelationId] = el.OriginalUrl
		res = append(res, CustomIdSet{CorrelationId: el.CorrelationId, ShortUrl: el.CorrelationId})
	}

	return res, nil
}
