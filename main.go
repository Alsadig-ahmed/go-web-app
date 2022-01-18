package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	//static folder
	app.Static("/", "./public")

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	app.Get("/h", func(c *fiber.Ctx) error {
		return c.SendFile("index.html")
	})

	// 404 page
	app.Get("*", func(c *fiber.Ctx) error {
		return c.SendFile("404.html")
	})

	port := os.Getenv("PORT")
	app.Listen(":" + port)
}
