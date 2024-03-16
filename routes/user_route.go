package route

import (
	"github.com/gofiber/fiber/v2"
	controller "github.com/jayesh.karande.dev/fibreMongo/controllers"
)

func SetRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		// Variable is now immutable
		return c.Send([]byte("Hellow from roots"))

	})
	app.Get("/users", controller.GetUsers)

}
