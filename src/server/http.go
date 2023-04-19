package server

import (
	"fmt"

	fiber "github.com/gofiber/fiber/v2"
)

type (
	HTTP struct {
		app  *fiber.App
		port int
	}
)

const (
	// TODO:
	// 	- This value should come from a config
	port = 9001
)

func NewHTTPServer() API {
	return &HTTP{
		app:  fiber.New(),
		port: port,
	}
}
func (h *HTTP) Start() error {
	return h.app.Listen(fmt.Sprintf(":%d", h.port))
}

func (h *HTTP) Close() error {
	return h.app.Shutdown()
}
