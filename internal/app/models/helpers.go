package models

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"time"
)

func SetCookie() (fiber.Cookie, error) {
	id := uuid.New().String() // crypto..
	return fiber.Cookie{
		Name:     "user",
		Value:    id,
		Secure:   true,
		Expires:  time.Now().Add(24 * 356 * time.Hour),
		HTTPOnly: true,
	}, nil
}

func ReadCookie(c *fiber.Ctx) string {
	// crypto
	return c.Cookies("value", "")
}
