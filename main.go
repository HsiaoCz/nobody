package main

import (
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	file, _ := os.OpenFile("./nobody.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	defer file.Close()
	app.Use(logger.New(logger.Config{
		Next:         nil,
		Format:       "[${time}] ${status} - ${latency} ${method} ${path}\n",
		TimeFormat:   "2006/01/02 15:04:05",
		TimeZone:     "Local",
		TimeInterval: 500 * time.Millisecond,
		Output:       io.MultiWriter(file, os.Stdout),
	}))
	app.Get("/hello", Hello)
	app.Listen("127.0.0.1:3001")
}

func Hello(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).SendString("Welcome to my show!")
}
