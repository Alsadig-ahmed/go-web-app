package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

type Message struct {
	Body string
}

func main() {
	app := fiber.New()

	//static folder
	app.Static("/", "./public")

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	// working with json
	msg := Message{"your request has been processed successfully"}

	app.Get("/api/request", func(c *fiber.Ctx) error {
		return c.JSON(msg)
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
