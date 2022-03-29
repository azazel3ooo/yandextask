package models

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"time"
)

func (s *Server) SetCookie() (fiber.Cookie, error) {
	value := map[string]string{
		"id": uuid.New().String(),
	}
	encoded, err := s.Cookie.Encode("user", value)
	if err == nil {
		return fiber.Cookie{
			Name:    "user",
			Value:   encoded,
			Expires: time.Now().Add(24 * 356 * time.Hour),
		}, nil
	}
	return fiber.Cookie{}, err
}

func (s *Server) ReadCookie(c *fiber.Ctx) string {
	if cookie := c.Cookies("user"); cookie != "" {
		value := make(map[string]string)
		if err := s.Cookie.Decode("user", cookie, &value); err == nil {
			return value["id"]
		}
	}
	return ""
}
