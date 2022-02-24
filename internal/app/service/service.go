package service

import (
	"encoding/json"
	"github.com/azazel3ooo/yandextask/internal/app/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"log"
	"net/http"
	"net/url"
)

func StartService() {
	config := models.InitConfig()
	app := fiber.New()

	app.Get("/:id", Getter)
	app.Post("/", Setter)
	app.Post("/api/shorten", JSONSetter)

	log.Fatal(app.Listen(config.Addr))
}

func Getter(c *fiber.Ctx) error {
	id := c.Params("id")
	if _, err := uuid.Parse(id); err != nil {
		return c.Status(http.StatusBadRequest).SendString("Невалидный id")
	}

	s, err := models.Store.Get(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}
	c.Location(s)
	return c.SendStatus(http.StatusTemporaryRedirect)
}

func Setter(c *fiber.Ctx) error {
	body := c.Body()
	u, err := url.ParseRequestURI(string(body))
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Невалидный URL")
	}

	return c.Status(http.StatusCreated).SendString(models.Store.Set(u.String()))
}

func JSONSetter(c *fiber.Ctx) error {
	body := c.Body()
	var req models.Request
	err := json.Unmarshal(body, &req)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid json")
	}
	_, err = url.ParseRequestURI(req.Url)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid URL")
	}

	return c.Status(http.StatusCreated).JSON(models.Response{
		Result: models.Store.Set(req.Url),
	})
}
