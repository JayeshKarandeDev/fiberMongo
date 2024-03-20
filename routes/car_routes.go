package route

import (
	"github.com/gofiber/fiber/v2"
	controller "github.com/jayesh.karande.dev/fibreMongo/controllers"
)

func carRoutes(app *fiber.App) {
	app.Get("/cars", controller.GetCars)
	app.Get("/car", controller.GetCar)

}
