package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/theHinneh/go-rest-api/database"
	"github.com/theHinneh/go-rest-api/todo"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/todo", todo.GetAllTodo)
	app.Get("/api/v1/todo/:id", todo.GetTodo)
	app.Post("/api/v1/todo", todo.NewTodo)
	app.Patch("/api/v1/todo/:id", todo.UpdateTodo)
	app.Delete("/api/v1/todo/:id", todo.DeleteTodo)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Databse opened")

	database.DBConn.AutoMigrate(&todo.Todo{})
}

func main() {
	app := fiber.New()
	initDatabase()
	// defer database.DB

	setupRoutes(app)

	app.Listen(":3000")
}
