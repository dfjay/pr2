package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http/httptest"
	"testing"
)

func TestHelloEndpoint(t *testing.T) {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world")
	})

	resp, err := app.Test(httptest.NewRequest("GET", "/", nil))
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello world", string(body))

	resp, err = app.Test(httptest.NewRequest("POST", "/", nil))
	assert.Nil(t, err)
	assert.Equal(t, 405, resp.StatusCode)
}
