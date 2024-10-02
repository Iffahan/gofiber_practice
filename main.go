package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Iffahan/gofiber_practice/models"
	"github.com/Iffahan/gofiber_practice/routers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors" // Import the CORS middleware
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

	// Auto-migrate the User model (create table if it doesn't exist)
	db.AutoMigrate(&models.User{})

	// Create a new Fiber instance
	app := fiber.New()

	// Enable CORS middleware with configuration
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",                    // Allow your Nuxt.js frontend (usually runs on localhost:3000)
		AllowMethods: "GET,POST,PUT,DELETE",                      // Allow methods
		AllowHeaders: "Origin,Content-Type,Accept,Authorization", // Allow necessary headers
	}))

	// Register routes
	routers.GetUsers(app, db)
	routers.RegisterUser(app, db) // Register the new routes
	routers.LoginUser(app, db)    // Register the login route
	routers.MyProfile(app, db)

	// Start the app on PORT 4000
	fmt.Println("App is running on http://localhost:4000")
	app.Listen(":4000")
}
