package models

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
