package router

import (
	"io"
	"os"
	"time"

	api "github.com/HsiaoCz/nobody/app/v1"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Route(r *fiber.App) {
	app := r.Group("/app")
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
	{
		v1 := app.Group("/v1")
		{
			auther := v1.Group("/auther")
			{
				auther.Get("/{id}", api.GetUserByID)
			}
		}
	}
}
