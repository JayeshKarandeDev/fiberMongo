package controller

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jayesh.karande.dev/fibreMongo/configs"
	"github.com/jayesh.karande.dev/fibreMongo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var carsCollection *mongo.Collection = configs.GetCollection(configs.DB, "cars")

func GetCars(c *fiber.Ctx) error {
	fmt.Println("Should return all cars")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var cars []models.Cars
	defer cancel()
	results, err := carsCollection.Find(ctx, bson.D{})
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	for results.Next(ctx) {
		var singleCar models.Cars
		if err = results.Decode(&singleCar); err != nil {
			return c.Status(http.StatusInternalServerError).SendString("Error decoding Car")
		}
		cars = append(cars, singleCar)
	}
	return c.JSON(cars)
	// return c.Send([]byte("cars list"))
}

func GetCar(c *fiber.Ctx) error {
	fmt.Println("Should return all cars")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var cars []models.Cars
	defer cancel()
	carName := c.Query("name")
	results, err := carsCollection.Find(ctx, bson.D{{"name", carName}})
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	for results.Next(ctx) {
		var singleCar models.Cars
		if err = results.Decode(&singleCar); err != nil {
			return c.Status(http.StatusInternalServerError).SendString("Error decoding Car")
		}
		cars = append(cars, singleCar)
	}
	return c.JSON(cars)
	// return c.Send([]byte("cars list"))
}
