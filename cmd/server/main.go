package main

import (
	"go_booking/cmd/config"
	"go_booking/internal/db"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadEnv()

	db.Connect()
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world GO booking")
	})

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	log.Println("Starting Fiber server on port", port)
	err := app.Listen(":8080")
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
