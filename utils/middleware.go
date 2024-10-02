package utils

import (
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// Protected is a middleware that checks if the user has a valid token
func Protected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Extract the token from the Authorization header
		tokenString := c.Get("Authorization")
		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing or invalid token"})
		}

		// In case the token has the 'Bearer ' prefix, you can choose to skip removing it or keep it clean
		// tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		// Parse the JWT token
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		// Debugging the token
		if err != nil {
			log.Printf("Error parsing token: %v\n", err)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}
		if !token.Valid {
			log.Println("Invalid token")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}

		// Store the claims in Locals for access in other handlers
		log.Printf("Token parsed successfully. Email: %s\n", claims.Email)
		c.Locals("user_claims", claims)

		// Proceed with the request
		return c.Next()
	}
}
