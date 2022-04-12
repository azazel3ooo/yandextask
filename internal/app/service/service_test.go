package service

import (
	"bytes"
	"github.com/azazel3ooo/yandextask/internal/app/models"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestSetter(t *testing.T) {
	type set struct {
		description  string
		route        string
		expectedCode int
		url          string
	}

	var (
		store models.Storage
		cfg   models.Config
	)

	store.Init(cfg)
	tempApp := fiber.New()
	s := models.NewServer(&store, cfg, tempApp)
	s.App.Post("/", s.Setter)

	tests := []set{
		{
			description:  "get HTTP status 201",
			route:        "/",
			expectedCode: http.StatusCreated,
			url:          "https://music.yandex.ru/home",
		},
		{
			description:  "get HTTP status 400 with invalid url",
			route:        "/",
			expectedCode: http.StatusBadRequest,
			url:          "Q_q",
		},
		{
			description:  "get HTTP status 404 with unknown route",
			route:        "/crash_me",
			expectedCode: http.StatusNotFound,
			url:          "https://yandex.ru",
		},
	}

	for _, test := range tests {
		b := bytes.NewBuffer([]byte(test.url))

		req := httptest.NewRequest(http.MethodPost, test.route, b)

		resp, _ := s.App.Test(req, 5)
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
		err := resp.Body.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}
}

func TestJSONSetter(t *testing.T) {
	type set struct {
		description  string
		route        string
		expectedCode int
		json         string
		error        string
	}

	var (
		store models.Storage
		cfg   models.Config
	)

	store.Init(cfg)
	tempApp := fiber.New()
	s := models.NewServer(&store, cfg, tempApp)
	s.App.Post("/api/shorten", s.JSONSetter)

	tests := []set{
		{
			description:  "get HTTP status 201",
			route:        "/api/shorten",
			expectedCode: http.StatusCreated,
			json:         "{\"url\": \"https://music.yandex.ru/home\"}",
		},
		{
			description:  "get HTTP status 400 with invalid url",
			route:        "/api/shorten",
			expectedCode: http.StatusBadRequest,
			json:         "{\"url\": \"<some_url>\"}",
			error:        "Invalid URL",
		},
		{
			description:  "get HTTP status 400 with invalid json",
			route:        "/api/shorten",
			expectedCode: http.StatusBadRequest,
			json:         "{\"url\": \"<some_url>\"",
			error:        "Invalid json",
		},
	}

	for _, test := range tests {
		b := bytes.NewBuffer([]byte(test.json))
		req := httptest.NewRequest(http.MethodPost, test.route, b)

		resp, _ := s.App.Test(req, 5)
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
		if test.expectedCode != http.StatusCreated {
			e, _ := io.ReadAll(resp.Body)
			assert.Equalf(t, test.error, string(e), test.description)
		}

		err := resp.Body.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}
}

func TestGetter(t *testing.T) {
	type set struct {
		description  string
		route        string
		expectedCode int
		url          string
	}

	var (
		store models.Storage
		cfg   models.Config
	)

	store.Init(cfg)
	tempApp := fiber.New()
	s := models.NewServer(&store, cfg, tempApp)
	s.App.Get("/:id", s.Getter)

	id, _ := s.Storage.Set("https://yandex.ru", "")

	tests := []set{
		{
			description:  "get success redirect 307",
			route:        "/" + id,
			expectedCode: http.StatusTemporaryRedirect,
			url:          "https://yandex.ru",
		},
		{
			description:  "get HTTP status 400 with invalid id",
			route:        "/a0dfasfa",
			expectedCode: http.StatusBadRequest,
			url:          "Q_q",
		},
		{
			description:  "get HTTP status 400 with unknown id",
			route:        "/" + uuid.New().String(),
			expectedCode: http.StatusBadRequest,
			url:          "https://yandex.ru",
		},
	}

	for _, test := range tests {

		req := httptest.NewRequest("GET", test.route, nil)

		resp, _ := s.App.Test(req, -1)

		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
		assert.Equalf(t, "text/plain; charset=utf-8", resp.Header.Get("Content-type"), test.description)
		if resp.StatusCode == http.StatusTemporaryRedirect {
			assert.Equalf(t, test.url, resp.Header.Get("Location"), test.description)
		}

		err := resp.Body.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}
}
