package route

import "github.com/gofiber/fiber/v2"

func SetRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		// Variable is now immutable
		return c.Send([]byte("Hellow from roots"))

	})
	userRoutes(app)
	carRoutes(app)
}
