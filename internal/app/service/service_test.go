package service

import (
	"bytes"
	"github.com/azazel3ooo/yandextask/internal/app/models"
	"github.com/google/uuid"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestSetter(t *testing.T) {
	t.Parallel()
	type set struct {
		description  string
		route        string
		expectedCode int
		url          string
	}

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

	app := fiber.New()
	app.Post("/", Setter)

	for _, test := range tests {
		b := bytes.NewBuffer([]byte(test.url))

		req := httptest.NewRequest(http.MethodPost, test.route, b)

		resp, _ := app.Test(req, 5)
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
		resp.Body.Close()
	}
}

func TestJSONSetter(t *testing.T) {
	t.Parallel()
	type set struct {
		description  string
		route        string
		expectedCode int
		json         string
		error        string
	}

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

	app := fiber.New()
	app.Post("/api/shorten", JSONSetter)

	for _, test := range tests {
		b := bytes.NewBuffer([]byte(test.json))
		req := httptest.NewRequest(http.MethodPost, test.route, b)

		resp, _ := app.Test(req, 5)
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
		if test.expectedCode != http.StatusCreated {
			e, _ := io.ReadAll(resp.Body)
			assert.Equalf(t, test.error, string(e), test.description)
		}
		resp.Body.Close()
	}
}

func TestGetter(t *testing.T) {
	t.Parallel()
	type set struct {
		description  string
		route        string
		expectedCode int
		url          string
	}

	tests := []set{
		{
			description:  "get success redirect 307",
			route:        "/" + models.Store.Set("https://yandex.ru"),
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

	//// Почему не работает?-_-
	//for _, test := range tests {
	//	if test.expectedCode == 307 {
	//		test.route = "/" + store.Set(test.url)
	//		log.Println(test.route)
	//	}
	//}

	app := fiber.New()
	app.Get("/:id", Getter)

	for _, test := range tests {

		req := httptest.NewRequest("GET", test.route, nil)

		resp, _ := app.Test(req, -1)

		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
		assert.Equalf(t, "text/plain; charset=utf-8", resp.Header.Get("Content-type"), test.description)
		if resp.StatusCode == http.StatusTemporaryRedirect {
			assert.Equalf(t, test.url, resp.Header.Get("Location"), test.description)
		}
		resp.Body.Close()
	}
}
