package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Iffahan/gofiber_practice/models"
	"github.com/Iffahan/gofiber_practice/routers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/Iffahan/gofiber_practice/docs" // Import generated docs
)

//	@title			Go Fiber Practice API
//	@version		1.0
//	@description	API documentation for Go Fiber practice project.
//	@host			localhost:3000
//	@BasePath		/
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

	// Auto-migrate the User model (create table if it doesn't exist)
	db.AutoMigrate(&models.User{})

	// Create a new Fiber instance
	app := fiber.New()

	// Setup Swagger docs route
	app.Get("/swagger/*", swagger.HandlerDefault) // Serve Swagger UI

	// Register routes
	routers.CreateUser(app, db)
	routers.GetUsers(app, db)

	// Start the app on PORT 3000
	fmt.Println("App is running on http://localhost:3000")
	app.Listen(":3000")
}
