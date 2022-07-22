package httptransport

import (
	"github.com/azazel3ooo/yandextask/internal/logic"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/azazel3ooo/yandextask/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/ugorji/go/codec"
)

// Getter - метод сервера, который предназначен для редиректа по короткой ссылку
func (s *Server) Getter(c *fiber.Ctx) error {
	id := c.Params("id")

	fullURL, err := logic.GetURL(id, s.Storage)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	if fullURL == "deleted" {
		return c.SendStatus(http.StatusGone)
	}
	c.Set("Location", fullURL)

	ck := logic.ReadCookie(c)
	tmp, _ := logic.SetCookie()
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

// Setter - метод сервера, который осуществляет создание новой короткой ссылки. Тело запроса - string
func (s *Server) Setter(c *fiber.Ctx) error {
	body := c.Body()

	ck := logic.ReadCookie(c)
	tmp, uid := logic.SetCookie()
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

	res, err := logic.SetURL(string(body), uid, s.Storage, s.Cfg.FileStoragePath, s.Cfg.URLBase)
	if err != nil {
		code, _ := strconv.Atoi(res)
		return c.Status(code).SendString(err.Error())
	}

	return c.Status(http.StatusCreated).SendString(res)
}

// JSONSetter - метод сервера, который осуществляет создание новой короткой ссылки. Тело запроса - json
func (s *Server) JSONSetter(c *fiber.Ctx) error {
	var req models.Request
	body := c.Body()
	handle := new(codec.JsonHandle)
	decoder := codec.NewDecoderBytes(body, handle)
	if err := decoder.Decode(&req); err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid json")
	}

	ck := logic.ReadCookie(c)
	tmp, uid := logic.SetCookie()
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

	// усложнилось обработка ошибок (для выбора правильного формата ответа), но в целом стало приятнее смотреть на хендлеры
	res, err := logic.SetURL(req.Addr, uid, s.Storage, s.Cfg.FileStoragePath, s.Cfg.URLBase)
	if err != nil {
		code, _ := strconv.Atoi(res)
		if code != http.StatusConflict {
			return c.Status(code).SendString(err.Error())
		}

		c.Set("Content-Type", "application/json")
		return c.Status(code).JSON(models.Response{Result: res})
	}

	c.Set("Content-Type", "application/json")
	return c.Status(http.StatusCreated).JSON(models.Response{
		Result: res,
	})
}

// UserUrlsGet - метод сервера, который возвращает полный список всех ссылок пользователя по его cookie
func (s *Server) UserUrlsGet(c *fiber.Ctx) error {
	ck := logic.ReadCookie(c)
	tmp, uid := logic.SetCookie()
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
	res, err := logic.UserUrlsGet(uid, s.Storage, s.Cfg.URLBase)
	if err != nil {
		log.Println(err)
		return c.SendStatus(http.StatusNoContent)
	}

	return c.Status(http.StatusOK).JSON(res)
}

// Ping - проверяет соединение с базой данных
func (s *Server) Ping(c *fiber.Ctx) error {
	if err := s.Storage.Ping(); err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.SendStatus(http.StatusOK)
}

// SetMany - метод сервера, который осуществляет множественное добалвение коротких ссылок с кастомными id
func (s *Server) SetMany(c *fiber.Ctx) error {
	var req []models.CustomIDSet
	body := c.Body()
	handle := new(codec.JsonHandle)
	decoder := codec.NewDecoderBytes(body, handle)
	if err := decoder.Decode(&req); err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid json")
	}

	ck := logic.ReadCookie(c)
	tmp, uid := logic.SetCookie()
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

	res, err := logic.SetMany(req, uid, s.Storage, s.Cfg.URLBase)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	return c.Status(http.StatusCreated).JSON(res)
}

// AsyncDelete - метод сервера, осуществляющий передачу ссылок в канал, для дальнейшего ассинхронного удаления
func (s *Server) AsyncDelete(c *fiber.Ctx) error {
	var (
		ids []string
	)

	bodyString := string(c.Body())
	bodyString = strings.ReplaceAll(bodyString, "[", "")
	bodyString = strings.ReplaceAll(bodyString, "]", "")
	bodyString = strings.ReplaceAll(bodyString, "\"", "")
	ids = strings.Split(bodyString, ",")

	ck := logic.ReadCookie(c)
	tmp, uid := logic.SetCookie()
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

	idsForDelete, err := logic.CheckUrlsForDelete(ids, uid, s.Storage)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	if len(idsForDelete) == 0 {
		return c.SendStatus(http.StatusNoContent)
	}

	s.ChanForDelete <- idsForDelete

	return c.SendStatus(http.StatusAccepted)
}

// GetStat - возвращает статистику по количеству записей, если Ip входит в WhiteList
func (s *Server) GetStat(c *fiber.Ctx) error {
	ip := c.GetRespHeader("X-Real-IP", "")
	if !logic.InWhiteList(ip, s.Cfg.Subnet) {
		return c.SendStatus(http.StatusForbidden)
	}

	users, urls, err := logic.GetStat(s.Storage)
	if err != nil {
		log.Println(err)
		return c.SendStatus(http.StatusInsufficientStorage)
	}
	res := models.StatInfo{Users: users, Urls: urls}
	return c.Status(http.StatusOK).JSON(res)
}
