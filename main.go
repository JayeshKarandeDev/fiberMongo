package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	route "github.com/jayesh.karande.dev/fibreMongo/routes"
)

func main() {
	app := fiber.New()
	// app.Get("/", func(c *fiber.Ctx) error {
	// 	// Variable is now immutable
	// 	return c.Send([]byte("Hellow from root"))

	// })
	route.SetRoutes(app)
	fmt.Println("App listening at localhost:4000")
	log.Fatal(app.Listen(":4000"))
}
