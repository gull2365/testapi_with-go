package main

import "github.com/gofiber/fiber/v2"

func data(c * fiber.Ctx) error{
	return c.JSON(fiber.Map{
		"massage" : "Hello World",
		"status" : 200,
	})
}

func main() {
	// make new api with fiber
	api := fiber.New()

	api.Get("/", func(c * fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	api.Get("/data",data)

	// start server
	api.Listen(":5000")
} 