package routers

import (
    "github.com/Iffahan/gofiber_practice/models"
    "github.com/gofiber/fiber/v2"
    "gorm.io/gorm"
)

// GetUsers returns all users
// @Summary Get all users
// @Description Retrieve all users from the database
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {array} models.User
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /users [get]
func GetUsers(app *fiber.App, db *gorm.DB) {
    app.Get("/users", func(c *fiber.Ctx) error {
        var users []models.User
        if result := db.Find(&users); result.Error != nil {
            return c.Status(500).JSON(map[string]string{"error": result.Error.Error()})
        }
        return c.JSON(users)
    })
}

// CreateUser creates a new user
// @Summary Create a new user
// @Description Add a new user to the database
// @Tags Users
// @Accept json
// @Produce json
// @Param user body models.User true "User data"
// @Success 201 {object} models.User
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /users [post]
func CreateUser(app *fiber.App, db *gorm.DB) {
    app.Post("/users", func(c *fiber.Ctx) error {
        user := new(models.User)
        if err := c.BodyParser(user); err != nil {
            return c.Status(400).JSON(map[string]string{"error": err.Error()})
        }
        if result := db.Create(&user); result.Error != nil {
            return c.Status(500).JSON(map[string]string{"error": result.Error.Error()})
        }
        return c.Status(201).JSON(user)
    })
}
