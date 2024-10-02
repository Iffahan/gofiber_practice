package routers

import (
	"github.com/Iffahan/gofiber_practice/models"
	"github.com/Iffahan/gofiber_practice/utils"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Register a new user
func RegisterUser(app *fiber.App, db *gorm.DB) {
	app.Post("/register", func(c *fiber.Ctx) error {
		user := new(models.User)

		// Parse the request body
		if err := c.BodyParser(user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse request"})
		}

		// Hash the password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not hash password"})
		}
		user.Password = string(hashedPassword)

		// Save the user to the database
		if err := db.Create(&user).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create user"})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User registered successfully"})
	})
}

// Login user and return JWT token
func LoginUser(app *fiber.App, db *gorm.DB) {
	app.Post("/login", func(c *fiber.Ctx) error {
		user := new(models.User)
		login := new(struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		})

		// Parse the request body
		if err := c.BodyParser(login); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse request"})
		}

		// Find the user in the database
		if err := db.Where("email = ?", login.Email).First(&user).Error; err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid email or password"})
		}

		// Compare the provided password with the stored hashed password
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)); err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid email or password"})
		}

		// Generate a JWT token
		token, err := utils.GenerateToken(user.Email)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not generate token"})
		}

		return c.JSON(fiber.Map{"token": token})
	})
}
