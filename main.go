package main

import (
 "fmt"
 "github.com/gofiber/fiber/v2"
)

func main() {
 fmt.Println("hello world")

 // fiber instance
 app := fiber.New()

 // routes 
 app.Get("/", func(c *fiber.Ctx) error {
  return c.SendString("hello world 🌈")
 })
 
// app listening at PORT: 3000
 app.Listen(":3000")
}