package main

import (
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/Iffahan/gofiber_practice/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Build the DSN from environment variables
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	// Connect to PostgreSQL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Drop all tables in the database
	fmt.Println("Dropping all tables...")
	if err := db.Migrator().DropTable(&models.User{}); err != nil {
		log.Fatalf("Error dropping table: %v", err)
	}
	// Repeat for other models as necessary
	// Add other model drops here if you have more than User
	// e.g., db.Migrator().DropTable(&models.OtherModel{})

	// Recreate all tables using AutoMigrate
	fmt.Println("Running database migrations...")
	modelsList := []interface{}{
		&models.User{},
		// Add other models here as needed
	}

	for _, model := range modelsList {
		if err := db.AutoMigrate(model); err != nil {
			log.Fatalf("Error during migration for model %s: %v", reflect.TypeOf(model).Elem().Name(), err)
		}
	}

	fmt.Println("Database initialization completed successfully!")
}
