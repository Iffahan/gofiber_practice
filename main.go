// main.go
package main

import (
	"fmt"
	"log"

	"github.com/Iffahan/gofiber_practice/models"
	"github.com/Iffahan/gofiber_practice/routers"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Database connection string
	dsn := "host=localhost user=admin password=admin1234 dbname=gofiber port=5432 sslmode=disable"

	// Connect to PostgreSQL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Auto migrate the User model (create table if it doesn't exist)
	db.AutoMigrate(&models.User{})

	// Create a new Fiber instance
	app := fiber.New()

	// Setup routes for users
	routers.SetupUserRoutes(app, db)

	// Start the app on PORT 3000
	fmt.Println("App is running on http://localhost:3000")
	app.Listen(":3000")
}
