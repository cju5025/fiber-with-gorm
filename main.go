package main

import (
	"fmt"
	"github.com/cju5025/fiber-with-gorm/book"
	"github.com/cju5025/fiber-with-gorm/database"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connection successfully opened.")

	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()

	initDatabase()
	// defer database.DBConn.Close()

	setupRoutes(app)

	app.Listen(":3000")
}
