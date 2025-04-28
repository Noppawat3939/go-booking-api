package db

import (
	"fmt"
	"go_booking/internal/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok",
		host, user, password, dbname, port,
	)

	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed connect database:", err)
	}

	err = models.MigrateDB(db)

	if err != nil {
		log.Fatalf("Failed migrate database %v", err)
	}

	fmt.Println("✅ Connected database successfully")
	DB = db
	return nil
}
