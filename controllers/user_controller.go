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

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "user")

func GetUsers(c *fiber.Ctx) error {
	fmt.Println("Should return all users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var users []models.User
	defer cancel()
	results, err := userCollection.Find(ctx, bson.D{})
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	for results.Next(ctx) {
		var singleUser models.User
		if err = results.Decode(&singleUser); err != nil {
			return c.Status(http.StatusInternalServerError).SendString("Error decoding user")
		}
		users = append(users, singleUser)
	}
	return c.JSON(users)
	// return c.Send([]byte("users list"))
}
