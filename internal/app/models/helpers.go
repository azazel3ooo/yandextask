package models

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"time"
)

func SetCookie() (*fiber.Cookie, string) {
	id := uuid.New().String() // crypto..
	ck := new(fiber.Cookie)
	ck.Name = "user"
	ck.Value = id
	ck.Expires = time.Now().Add(24 * 356 * time.Hour)
	return ck, id
	//return fiber.Cookie{
	//	Name:     id,
	//	Secure:   true,
	//	Expires:  time.Now().Add(24 * 356 * time.Hour),
	//	HTTPOnly: true,
	//}, nil
}

func ReadCookie(c *fiber.Ctx) string {
	// crypto
	return c.Cookies("user")
}
