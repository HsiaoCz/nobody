package main

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/hello", Hello)
	app.Listen("127.0.0.1:3001")
}

func Hello(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).SendString("Welcome to my show!")
}
