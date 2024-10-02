package routers

import (
	"github.com/Iffahan/gofiber_practice/models"
	"github.com/Iffahan/gofiber_practice/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GetUsers retrieves all users from the database
func GetUsers(app *fiber.App, db *gorm.DB) {
	app.Get("/users", utils.Protected(), func(c *fiber.Ctx) error {
		var users []models.User
		if err := db.Find(&users).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not retrieve users"})
		}
		return c.JSON(users)
	})
}

// MyProfile retrieves the profile of the currently authenticated user based on email from the token
func MyProfile(app *fiber.App, db *gorm.DB) {
	app.Get("/profile", utils.Protected(), func(c *fiber.Ctx) error {
		// Get the email from the token claims
		claims := c.Locals("user_claims").(*utils.Claims)
		email := claims.Email

		// Fetch user profile using the email from the database
		var user models.User
		if err := db.Where("email = ?", email).First(&user).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
		}

		// Return the user profile data
		return c.JSON(fiber.Map{
			"ID":    user.ID,
			"email": user.Email,
		})
	})
}
