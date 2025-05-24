package main

import (
	cfg "go_booking/cmd/config"
	"go_booking/internal/adapters/handler/router"
	"go_booking/internal/db"

	"log"
	"os"

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func init() {
	cfg.LoadEnv()
}

func main() {
	db, _ := db.ConnectDB()

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowMethods: "GET,POST,PATCH",
		AllowHeaders: "Content-Type,Authorization",
	}))

	app.Use(swagger.New(cfg.Swagger()))
	router.InitialRoutes(app, db)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	log.Println("Starting Fiber server on port", port)
	app.Listen(":8080")
}
