package main

import (
	"go_booking/cmd/config"
	"go_booking/internal/db"
	"go_booking/internal/routes"
	"log"
	"os"

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func init() {
	config.LoadEnv()
}

func main() {
	db, _ := db.ConnectDB()

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowMethods: "GET,POST,PATCH",
		AllowHeaders: "Content-Type",
	}))

	routes.NewRoutes(app, db)

	app.Use(swagger.New(config.Swagger()))
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	log.Println("Starting Fiber server on port", port)
	app.Listen(":8080")
}
