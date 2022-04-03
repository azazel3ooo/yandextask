package models

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"net/http"
	"net/url"
	"time"
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
	c.Cookie(&fiber.Cookie{
		Name:    "user",
		Value:   ReadCookie(c),
		Expires: time.Now().Add(24 * 356 * time.Hour),
	})
	return c.SendStatus(http.StatusTemporaryRedirect)
}

func (s *Server) Setter(c *fiber.Ctx) error {
	body := c.Body()
	u, err := url.ParseRequestURI(string(body))
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Невалидный URL")
	}

	ck := ReadCookie(c)
	tmp, uid := SetCookie()
	if ck == "" {
		c.Cookie(tmp)
	}
	id, err := s.Storage.Set(u.String(), s.Cfg.FileStoragePath)
	result := s.Cfg.URLBase + "/" + id
	if err != nil && id != "" {
		return c.Status(http.StatusConflict).JSON(Response{
			Result: result,
		})
	}
	if id == "" {
		return c.SendStatus(http.StatusInsufficientStorage)
	}

	s.Storage.UsersSet(uid, id)

	return c.Status(http.StatusCreated).SendString(result)
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

	ck := ReadCookie(c)
	tmp, uid := SetCookie()
	if ck == "" {
		c.Cookie(tmp)
	}

	id, err := s.Storage.Set(req.Addr, s.Cfg.FileStoragePath)
	result := s.Cfg.URLBase + "/" + id
	if err != nil && id != "" {
		return c.Status(http.StatusConflict).JSON(Response{
			Result: result,
		})
	}
	if id == "" {
		return c.SendStatus(http.StatusInsufficientStorage)
	}

	s.Storage.UsersSet(uid, id)

	c.Set("Content-Type", "application/json")
	return c.Status(http.StatusCreated).JSON(Response{
		Result: result,
	})
}

func (s *Server) UserUrlsGet(c *fiber.Ctx) error {
	ck := ReadCookie(c)
	if ck == "" {
		return c.SendStatus(http.StatusNoContent)
	}
	//log.Println(ck)
	ids, err := s.Storage.UsersGet(ck)
	if err != nil {
		return c.SendStatus(http.StatusNoContent)
	}
	res, _ := s.Storage.GetUrlsForUser(ids)
	for idx, el := range res {
		res[idx].Short = s.Cfg.URLBase + "/" + el.Short
		res[idx].Original = el.Original
	}
	c.Cookie(&fiber.Cookie{
		Name:    "user",
		Value:   ReadCookie(c),
		Expires: time.Now().Add(24 * 356 * time.Hour),
	})
	return c.Status(http.StatusOK).JSON(res)
}

func (s *Server) Ping(c *fiber.Ctx) error {
	if err := s.Storage.Ping(); err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.SendStatus(http.StatusOK)
}

func (s *Server) SetMany(c *fiber.Ctx) error {
	var req []CustomIDSet
	body := c.Body()
	err := json.Unmarshal(body, &req)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid json")
	}
	for _, el := range req {
		_, err = url.ParseRequestURI(el.OriginalURL)
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString("Invalid URL")
		}
	}

	ck := ReadCookie(c)
	tmp, uid := SetCookie()
	if ck == "" {
		c.Cookie(tmp)
	}

	res, _ := s.Storage.InsertMany(req)
	for idx, el := range res {
		res[idx].ShortURL = s.Cfg.URLBase + "/" + el.ShortURL
		s.Storage.UsersSet(uid, el.CorrelationID)
	}

	return c.Status(http.StatusCreated).JSON(res)
}
