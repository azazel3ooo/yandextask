package models

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/caarlos0/env/v6"
	"github.com/google/uuid"
)

func flagExist(flagName string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == flagName {
			found = true
		}
	})
	return found
}

func mergeConfigs(mainConf *Config, tmp Config) {
	if _, existEnv := os.LookupEnv("SERVER_ADDRESS"); !existEnv {
		if !flagExist("a") {
			mainConf.ServerAddress = tmp.ServerAddress
		}
	}
	if _, existEnv := os.LookupEnv("BASE_URL"); !existEnv {
		if !flagExist("b") {
			mainConf.URLBase = tmp.URLBase
		}
	}
	if _, existEnv := os.LookupEnv("FILE_STORAGE_PATH"); !existEnv {
		if !flagExist("f") {
			mainConf.FileStoragePath = tmp.FileStoragePath
		}
	}
	if _, existEnv := os.LookupEnv("DATABASE_DSN"); !existEnv {
		if !flagExist("d") {
			mainConf.DatabaseDsn = tmp.DatabaseDsn
		}
	}
	if _, existEnv := os.LookupEnv("ENABLE_HTTPS"); !existEnv {
		if !flagExist("s") {
			mainConf.EnableTLS = tmp.EnableTLS
		}
	}
}

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
	if c.Subnet == "" {
		flag.StringVar(&c.Subnet, "t", "", "Subnet")
	}
	if c.GAddress == "" {
		flag.StringVar(&c.GAddress, "g", "localhost:8082", "Grpc address")
	}

	if _, exist := os.LookupEnv("ENABLE_HTTPS"); !exist {
		flag.BoolVar(&c.EnableTLS, "s", false, "Enable https")
	}

	var configPath string
	if path, exist := os.LookupEnv("CONFIG"); !exist {
		flag.StringVar(&configPath, "c", "", "Use config file")
	} else {
		configPath = path
	}

	flag.Parse()

	if configPath != "" {
		var tmpConfig Config
		configFile, err := os.ReadFile(configPath)
		if err != nil {
			return err
		}

		err = json.Unmarshal(configFile, &tmpConfig)
		if err != nil {
			return err
		}

		mergeConfigs(c, tmpConfig)
	}

	return nil
}

func initData() data {
	return make(map[string]string)
}

func initUsers() users {
	return make(map[string][]string)
}

func uploadData(s *Storage, cfg Config) {
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
	s.Data = initData()
	if cfg.FileStoragePath != "" {
		if _, err := os.Stat(cfg.FileStoragePath); !os.IsNotExist(err) {
			uploadData(s, cfg)
		}
	}
	s.Users = initUsers()
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
		setToFile(id.String(), val, pth)
	}

	return id.String(), nil
}

func setToFile(k, v, pth string) {
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

// UsersGet - вовзращает ссылки по id пользователя
func (s *Storage) UsersGet(id string) ([]string, error) {
	urls, ok := s.Users[id]
	if !ok {
		return []string{}, nil
	}

	return urls, nil
}

// UsersSet - добавляет пользователю новую ссылку
func (s *Storage) UsersSet(id, url string) error {
	urls, ok := s.Users[id]
	if !ok {
		s.Users[id] = []string{url}
	} else {
		s.Users[id] = append(urls, url)
	}
	return nil
}

// GetUrlsForUser - возвращает все ссылки по массиву id
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

// Ping - говорит, что это не бд
func (s *Storage) Ping() error {
	return errors.New("i'm not db")
}

// InsertMany - добавляет несколько новых ссылок с кастомными id
func (s *Storage) InsertMany(m []CustomIDSet) ([]CustomIDSet, error) {
	var res []CustomIDSet
	for _, el := range m {
		s.Data[el.CorrelationID] = el.OriginalURL
		res = append(res, CustomIDSet{CorrelationID: el.CorrelationID, ShortURL: el.CorrelationID})
	}

	return res, nil
}

func (s *Storage) Delete(ids []string) error {

	return nil
}

func (s *Storage) GetUserStat() (int, error) {
	return len(s.Users), nil
}

func (s *Storage) GetUrlsStat() (int, error) {
	return len(s.Data), nil
}
