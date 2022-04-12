package models

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (s *Server) Getter(c *fiber.Ctx) error {
	id := c.Params("id")
	if _, err := uuid.Parse(id); err != nil {
		return c.Status(http.StatusBadRequest).SendString("Невалидный id")
	}
	fullURL, err := s.Storage.Get(id)
	if fullURL == "url" {
		return c.SendStatus(http.StatusGone)
	}
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}
	c.Set("Location", fullURL)

	ck := ReadCookie(c)
	tmp, _ := SetCookie()
	if ck == "" {
		c.Cookie(tmp)
	} else {
		c.Cookie(&fiber.Cookie{
			Name:    "user",
			Value:   ck,
			Expires: time.Now().Add(24 * 356 * time.Hour),
		})
	}

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
	} else {
		uid = ck
		c.Cookie(&fiber.Cookie{
			Name:    "user",
			Value:   ck,
			Expires: time.Now().Add(24 * 356 * time.Hour),
		})
	}

	id, err := s.Storage.Set(u.String(), s.Cfg.FileStoragePath)
	result := s.Cfg.URLBase + "/" + id
	if err != nil && id != "" {
		return c.Status(http.StatusConflict).SendString(result)
	}
	if id == "" {
		log.Println(err)
		return c.SendStatus(http.StatusInsufficientStorage)
	}

	err = s.Storage.UsersSet(uid, id)
	if err != nil {
		log.Println(err.Error())
	}

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
	} else {
		uid = ck
		c.Cookie(&fiber.Cookie{
			Name:    "user",
			Value:   ck,
			Expires: time.Now().Add(24 * 356 * time.Hour),
		})
	}

	id, err := s.Storage.Set(req.Addr, s.Cfg.FileStoragePath)
	result := s.Cfg.URLBase + "/" + id
	if err != nil && id != "" {
		return c.Status(http.StatusConflict).JSON(Response{
			Result: result,
		})
	}
	if id == "" {
		log.Println(err)
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
	tmp, uid := SetCookie()
	if ck == "" {
		c.Cookie(tmp)
	} else {
		uid = ck
		c.Cookie(&fiber.Cookie{
			Name:    "user",
			Value:   ck,
			Expires: time.Now().Add(24 * 356 * time.Hour),
		})
	}

	if ck == "" {
		return c.SendStatus(http.StatusNoContent)
	}
	ids, err := s.Storage.UsersGet(uid)
	if err != nil {
		return c.SendStatus(http.StatusNoContent)
	}
	res, _ := s.Storage.GetUrlsForUser(ids)
	for idx, el := range res {
		res[idx].Short = s.Cfg.URLBase + "/" + el.Short
		res[idx].Original = el.Original
	}

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
	} else {
		uid = ck
		c.Cookie(&fiber.Cookie{
			Name:    "user",
			Value:   ck,
			Expires: time.Now().Add(24 * 356 * time.Hour),
		})
	}

	res, _ := s.Storage.InsertMany(req)
	for idx, el := range res {
		res[idx].ShortURL = s.Cfg.URLBase + "/" + el.ShortURL
		s.Storage.UsersSet(uid, el.CorrelationID)
	}

	return c.Status(http.StatusCreated).JSON(res)
}

func (s *Server) AsyncDelete(c *fiber.Ctx) error {
	var (
		ids          []string
		idsForDelete []string
	)
	err := c.BodyParser(&ids)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("invalid body")
	}

	ck := ReadCookie(c)
	tmp, uid := SetCookie()
	if ck == "" {
		c.Cookie(tmp)
	} else {
		uid = ck
		c.Cookie(&fiber.Cookie{
			Name:    "user",
			Value:   ck,
			Expires: time.Now().Add(24 * 356 * time.Hour),
		})
	}

	urls, err := s.Storage.UsersGet(uid)
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	urlsMap := make(map[string]struct{})
	for _, id := range urls {
		urlsMap[id] = struct{}{}
	}

	for _, id := range ids {
		_, ok := urlsMap[id]
		if ok {
			idsForDelete = append(idsForDelete, id)
		}
	}

	if len(idsForDelete) == 0 {
		log.Println(urls)
		log.Println(ids)
		return c.SendStatus(http.StatusNoContent)
	}

	s.ChanForDelete <- idsForDelete

	return c.SendStatus(http.StatusAccepted)
}
