package router

import (
	api "github.com/HsiaoCz/nobody/app/v1"
	"github.com/gofiber/fiber/v2"
)

func Route(r *fiber.App) {
	app := r.Group("/app")
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
