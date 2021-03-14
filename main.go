package main

import (
	"github.com/cju5025/fiber-with-gorm/book"
	"github.com/cju5025/fiber-with-gorm/database"
	"github.com/gofiber/fiber/v2"
)

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World")
}

func setupRoutes(app *fiber.App) {
	app.Get("/books", book.GetBooks)
	app.Get("/books/:id", book.GetBook)
	app.Post("/books", book.NewBook)
	app.Delete("/books/:id", book.DeleteBook)
}

func main() {
	app := fiber.New()

	setupRoutes(app)

	app.Listen(":3000")
}
