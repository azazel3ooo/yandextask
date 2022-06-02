package server

import (
	"github.com/azazel3ooo/yandextask/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"log"
	"sync"
	"time"
)

func SetCookie() (*fiber.Cookie, string) {
	id := uuid.New().String() // crypto..
	ck := new(fiber.Cookie)
	ck.Name = "user"
	ck.Value = id
	ck.Expires = time.Now().Add(24 * 356 * time.Hour)
	return ck, id
}

func ReadCookie(c *fiber.Ctx) string {
	// crypto
	return c.Cookies("user")
}

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
