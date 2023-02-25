package main

import (
	"context"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joehann-9s/gtd-api/api/v1/routes"
	"github.com/joehann-9s/gtd-api/pkg/database"
	"github.com/joehann-9s/gtd-api/pkg/utils"

	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	//load env
	utils.LoadEnv()

	//
	port := os.Getenv("PORT")
	if port == "" {
		port = "7000"
	}

	app := fiber.New()
	app.Use(cors.New())

	// database
	client, err := database.LoadDB()
	if err != nil {
		panic(err)
	}

	// insert data
	coll := client.Database("testing").Collection("tasks")
	coll.InsertOne(context.TODO(), bson.D{{
		Key:   "name",
		Value: "coocking",
	}})
	//routes
	// import routes from userRoutes.go
	routes.UserRoutes(app)

	app.Listen(":" + port)
}
