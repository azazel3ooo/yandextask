package server

import (
	"log"
	"sync"
	"time"

	"github.com/azazel3ooo/yandextask/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// setCookie создает новую сессию
func setCookie() (*fiber.Cookie, string) {
	id := uuid.New().String() // crypto..
	ck := new(fiber.Cookie)
	ck.Name = "user"
	ck.Value = id
	ck.Expires = time.Now().Add(24 * 356 * time.Hour)
	return ck, id
}

// readCookie принимает на вход сессию и возвращает id пользователя
func readCookie(c *fiber.Ctx) string {
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
