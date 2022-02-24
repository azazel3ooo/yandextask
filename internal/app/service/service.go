package service

import (
	"encoding/json"
	"github.com/azazel3ooo/yandextask/internal/app/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/google/uuid"
	"log"
	"net/http"
	"net/url"
)

func StartService() {
	Conf := models.InitConfig()
	app := fiber.New()
	app.Use(recover.New(recover.Config{EnableStackTrace: true}))
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${latency} ${method} ${path} ${resBody}\n",
	}))
	app.Get("/:id", Getter)
	app.Post("/", Setter)
	app.Post("/api/shorten", JSONSetter)

	log.Fatal(app.Listen(Conf.Addr))
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
	c.Set("Location", s)
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
	_, err = url.ParseRequestURI(req.Addr)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid URL")
	}

	return c.Status(http.StatusCreated).JSON(models.Response{
		Result: models.Store.Set(req.Addr),
	})
}
