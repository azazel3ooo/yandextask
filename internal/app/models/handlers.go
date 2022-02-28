package models

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"net/http"
	"net/url"
)

func (s *Server) Getter(c *fiber.Ctx) error {
	id := c.Params("id")
	if _, err := uuid.Parse(id); err != nil {
		return c.Status(http.StatusBadRequest).SendString("Невалидный id")
	}

	fullURL, err := s.Storage.Get(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}
	c.Set("Location", fullURL)
	return c.SendStatus(http.StatusTemporaryRedirect)
}

func (s *Server) Setter(c *fiber.Ctx) error {
	body := c.Body()
	u, err := url.ParseRequestURI(string(body))
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Невалидный URL")
	}

	host := fmt.Sprintf("%s:%d", s.Cfg.URLBase, s.Cfg.ServerAddress)
	return c.Status(http.StatusCreated).SendString(host + "/" + s.Storage.Set(u.String(), s.Cfg.FileStoragePath))
}

func (s *Server) JSONSetter(c *fiber.Ctx) error {
	body := c.Body()
	var req Request
	err := json.Unmarshal(body, &req)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid json")
	}
	_, err = url.ParseRequestURI(req.Addr)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid URL")
	}

	host := fmt.Sprintf("%s:%d", s.Cfg.URLBase, s.Cfg.ServerAddress)
	c.Set("Content-Type", "application/json")
	return c.Status(http.StatusCreated).JSON(Response{
		Result: host + "/" + s.Storage.Set(req.Addr, s.Cfg.FileStoragePath),
	})
}
