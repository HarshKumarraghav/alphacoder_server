package main

import (
	"alphacoder/pkg/configuration"
	"context"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	app := fiber.New()
	def := cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowCredentials: true,
	}
	app.Use(cors.New(def))
	godotenv.Load()
	config := configuration.FromEnv()
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.MongoURI))
	if err != nil {
		log.Panic(err)
	}
	db := client.Database("botsfusion_chatbotbuilder")
	log.Panic(app.Listen(":" + os.Getenv("PORT")))

}
