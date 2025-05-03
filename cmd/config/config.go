package config

import (
	"log"

	"github.com/gofiber/contrib/swagger"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Failed load env file")
	}

	log.Println("Loaded env is success")
}

func Swagger() swagger.Config {
	log.Println("✅ Intialized swagger config")

	return swagger.Config{
		BasePath: "/",
		FilePath: "./swagger.json",
		Path:     "swagger",
		Title:    "Swagger Booking API Docs",
	}
}
