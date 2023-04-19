package server

import (
	fiber "github.com/gofiber/fiber/v2"
)

type (
	HTTP struct {
		app fiber.App
	}
)

const (
	// TODO:
	// 	- This value should come from a config
	port = "9001"
)

func NewHTTPServer() API {
	return &HTTP{}
}
func (h *HTTP) Start() error {
	return nil
}

func (h *HTTP) Close() error {
	return nil
}
