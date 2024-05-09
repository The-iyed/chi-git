package api

import (
    "github.com/gofiber/fiber/v2"
	"go-chi/database"
)

func Run() {
    app := fiber.New()
    app.Get("/", func(c *fiber.Ctx) error {
        return c.JSON(&fiber.Map{"data": "Hello from server !"})
    })
    database.Connect()
    routes.UserRoute(app)
    app.Listen(":5003")
}
