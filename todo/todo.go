package todo

import (
	"github.com/gofiber/fiber/v2"
	"github.com/theHinneh/go-rest-api/database"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Location    string `json:"location"`
	Date        string `json:"date"`
	Time        string `json:"time"`
}

func GetAllTodo(c *fiber.Ctx) error {
	db := database.DBConn
	var todo []Todo
	db.Find(&todo)
	return c.JSON(todo)
}

func GetTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var todo Todo
	db.Find(&todo, id)
	return c.JSON(todo)
}

func NewTodo(c *fiber.Ctx) error {
	db := database.DBConn

	todo := new(Todo)
	if err := c.BodyParser(todo); err != nil {
		return c.Status(500).SendString("Ann error occured")
	}

	db.Create(&todo)
	return c.JSON(todo)
}

func DeleteTodo(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var todo Todo

	db.First(&todo, id)
	if todo.Title == "" {
		return c.Status(500).SendString("No Todo found")
	}
	db.Delete(&todo)
	return c.SendString("Todo deleted")
}

func UpdateTodo(c *fiber.Ctx) error {
	return c.SendString("Update a Todo")
}
