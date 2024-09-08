// routers/user.go
package routers

import (
	"github.com/Iffahan/gofiber_practice/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// SetupUserRoutes sets up the routes for user-related endpoints
func SetupUserRoutes(app *fiber.App, db *gorm.DB) {
    app.Get("/users", func(c *fiber.Ctx) error {
        var users []models.User
        if result := db.Find(&users); result.Error != nil {
            return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
        }
        return c.JSON(users)
    })

    app.Post("/users", func(c *fiber.Ctx) error {
        user := new(models.User)
        if err := c.BodyParser(user); err != nil {
            return c.Status(400).JSON(fiber.Map{"error": err.Error()})
        }
        if result := db.Create(&user); result.Error != nil {
            return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
        }
        return c.JSON(user)
    })
}
