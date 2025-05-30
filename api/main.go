package main

import (
	"fmt"
	"log"
	"os"

	"github.com/adix/shortapi/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func setupRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveUrl)
	app.Post("/api/v1", routes.ShortenUrl)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}
	app := fiber.New()
	app.Use(logger.New())

	setupRoutes(app)

	fmt.Println("Server is starting on port: %d", os.Getenv("APP_PORT"))
	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}
