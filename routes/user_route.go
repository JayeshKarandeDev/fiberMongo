package route

import (
	"github.com/gofiber/fiber/v2"
	controller "github.com/jayesh.karande.dev/fibreMongo/controllers"
)

func userRoutes(app *fiber.App) {
	app.Get("/users", controller.GetUsers)
}
