package logic

import (
	"errors"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/azazel3ooo/yandextask/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// SetCookie создает новую сессию
func SetCookie() (*fiber.Cookie, string) {
	id := uuid.New().String() // crypto..
	ck := new(fiber.Cookie)
	ck.Name = "user"
	ck.Value = id
	ck.Expires = time.Now().Add(24 * 356 * time.Hour)
	return ck, id
}

// ReadCookie принимает на вход сессию и возвращает id пользователя
func ReadCookie(c *fiber.Ctx) string {
	// crypto
	return c.Cookies("user")
}

// FanIn вотчер, который ожидает массив id записей и вызывает асинхронное удаление
func FanIn(c chan []string, generalWt *sync.WaitGroup, storage models.Storable) {
	var wt sync.WaitGroup
	maxWorkers := 4
	goroutines := make(chan struct{}, maxWorkers)
	defer close(goroutines)

	for ids := range c {
		log.Println("FanIn get", ids)
		wt.Add(1)
		goroutines <- struct{}{}

		go deleteIds(&wt, ids, goroutines, storage)
	}
	wt.Wait()
	generalWt.Done()
}

func deleteIds(wt *sync.WaitGroup, ids []string, goroutines chan struct{}, s models.Storable) {
	err := s.Delete(ids)
	if err != nil {
		log.Println(err)
	}

	wt.Done()
	<-goroutines
}

// InWhiteList - проверяет входит ли Ip в доверенную сеть
func InWhiteList(ip, subnet string) bool {
	if subnet == "" {
		return false
	}
	return strings.HasPrefix(ip, subnet)
}

func GetUrl(id string, storage models.Storable) (string, error) {
	if _, err := uuid.Parse(id); err != nil {
		return "", errors.New("Невалидный id")
	}
	fullURL, err := storage.Get(id)
	if err != nil {
		return "", err
	}

	return fullURL, nil
}

func SetUrl(u, user string, storage models.Storable, pathToFileStorage, baseURL string) (string, error) {
	reqUrl, err := url.ParseRequestURI(u)
	if err != nil {
		return strconv.Itoa(http.StatusBadRequest), errors.New("Invalid URL")
	}

	id, err := storage.Set(reqUrl.String(), pathToFileStorage)
	result := baseURL + "/" + id
	if err != nil && id != "" {
		return strconv.Itoa(http.StatusConflict), errors.New(result)
	}
	if id == "" {
		log.Println(err)
		return strconv.Itoa(http.StatusInsufficientStorage), err
	}

	err = storage.UsersSet(user, id)
	if err != nil {
		log.Println(err)
	}

	return result, nil
}

func UserUrlsGet(user string, storage models.Storable, baseURL string) ([]models.UserResponse, error) {
	ids, err := storage.UsersGet(user)
	if err != nil {
		return nil, err
	}
	res, _ := storage.GetUrlsForUser(ids)
	for idx, el := range res {
		res[idx].Short = baseURL + "/" + el.Short
		res[idx].Original = el.Original
	}
	return res, nil
}

func SetMany(pairs []models.CustomIDSet, user string, storage models.Storable, baseURL string) ([]models.CustomIDSet, error) {
	for _, el := range pairs {
		_, err := url.ParseRequestURI(el.OriginalURL)
		if err != nil {
			return nil, errors.New("Invalid URL")
		}
	}

	res, _ := storage.InsertMany(pairs)
	for idx, el := range res {
		res[idx].ShortURL = baseURL + "/" + el.ShortURL
		storage.UsersSet(user, el.CorrelationID)
	}

	return res, nil
}

func CheckUrlsForDelete(ids []string, user string, storage models.Storable) ([]string, error) {
	var idsForDelete []string

	urls, err := storage.UsersGet(user)
	if err != nil {
		return nil, err
	}

	urlsMap := make(map[string]struct{})
	for _, id := range urls {
		urlsMap[id] = struct{}{}
	}

	for _, id := range ids {
		_, ok := urlsMap[id]
		if ok {
			idsForDelete = append(idsForDelete, id)
		}
	}

	return idsForDelete, nil
}

func GetStat(storage models.Storable) (int, int, error) {
	var (
		res models.StatInfo
		err error
	)
	res.Users, err = storage.GetUserStat()
	if err != nil {
		return 0, 0, err
	}
	res.Urls, err = storage.GetUrlsStat()
	if err != nil {
		return 0, 0, err
	}
	return res.Users, res.Urls, nil
}
